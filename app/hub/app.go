package hub

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	_staking "github.com/ironman0x7b2/sentinel-sdk/x/staking"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

const (
	appName        = "Sentinel Hub"
	DefaultKeyPass = "1234567890"
)

// nolint:gochecknoglobals
var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hubcli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.hubd")
)

type Hub struct {
	*baseapp.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	keyMain *csdk.KVStoreKey

	keyParams        *csdk.KVStoreKey
	keyAccount       *csdk.KVStoreKey
	keyFeeCollection *csdk.KVStoreKey
	keyStaking       *csdk.KVStoreKey
	keySlashing      *csdk.KVStoreKey
	keyDistribution  *csdk.KVStoreKey
	keyGov           *csdk.KVStoreKey
	keyMint          *csdk.KVStoreKey

	keyDeposit         *csdk.KVStoreKey
	keyVPNNode         *csdk.KVStoreKey
	keyVPNSession      *csdk.KVStoreKey
	keyVPNSubscription *csdk.KVStoreKey

	tkeyParams       *csdk.TransientStoreKey
	tkeyStaking      *csdk.TransientStoreKey
	tkeyDistribution *csdk.TransientStoreKey

	paramsKeeper        params.Keeper
	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	stakingKeeper       staking.Keeper
	slashingKeeper      slashing.Keeper
	distributionKeeper  distribution.Keeper
	govKeeper           gov.Keeper
	mintKeeper          mint.Keeper
	crisisKeeper        crisis.Keeper

	depositKeeper deposit.Keeper
	vpnKeeper     vpn.Keeper
}

func NewHub(logger log.Logger, db tmDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint,
	baseAppOptions ...func(*baseapp.BaseApp)) *Hub {

	cdc := MakeCodec()

	bApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	var app = &Hub{
		BaseApp:            bApp,
		cdc:                cdc,
		invCheckPeriod:     invCheckPeriod,
		keyParams:          csdk.NewKVStoreKey(params.StoreKey),
		keyMain:            csdk.NewKVStoreKey(baseapp.MainStoreKey),
		keyAccount:         csdk.NewKVStoreKey(auth.StoreKey),
		keyFeeCollection:   csdk.NewKVStoreKey(auth.FeeStoreKey),
		keyStaking:         csdk.NewKVStoreKey(staking.StoreKey),
		keySlashing:        csdk.NewKVStoreKey(slashing.StoreKey),
		keyDistribution:    csdk.NewKVStoreKey(distribution.StoreKey),
		keyGov:             csdk.NewKVStoreKey(gov.StoreKey),
		keyMint:            csdk.NewKVStoreKey(mint.StoreKey),
		keyDeposit:         csdk.NewKVStoreKey(deposit.StoreKey),
		keyVPNNode:         csdk.NewKVStoreKey(vpn.StoreKeyNode),
		keyVPNSession:      csdk.NewKVStoreKey(vpn.StoreKeySession),
		keyVPNSubscription: csdk.NewKVStoreKey(vpn.StoreKeySubscription),
		tkeyParams:         csdk.NewTransientStoreKey(params.TStoreKey),
		tkeyStaking:        csdk.NewTransientStoreKey(staking.TStoreKey),
		tkeyDistribution:   csdk.NewTransientStoreKey(distribution.TStoreKey),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc,
		app.keyParams,
		app.tkeyParams)
	app.accountKeeper = auth.NewAccountKeeper(app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace)
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(app.cdc,
		app.keyFeeCollection)
	stakingKeeper := staking.NewKeeper(app.cdc,
		app.keyStaking,
		app.tkeyStaking,
		app.bankKeeper,
		app.paramsKeeper.Subspace(staking.DefaultParamspace),
		staking.DefaultCodespace)
	app.slashingKeeper = slashing.NewKeeper(app.cdc,
		app.keySlashing,
		&stakingKeeper,
		app.paramsKeeper.Subspace(slashing.DefaultParamspace),
		slashing.DefaultCodespace)
	app.distributionKeeper = distribution.NewKeeper(app.cdc,
		app.keyDistribution,
		app.paramsKeeper.Subspace(distribution.DefaultParamspace),
		app.bankKeeper,
		&stakingKeeper,
		app.feeCollectionKeeper,
		distribution.DefaultCodespace)
	app.govKeeper = gov.NewKeeper(app.cdc,
		app.keyGov,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(gov.DefaultParamspace),
		app.bankKeeper,
		&stakingKeeper,
		gov.DefaultCodespace)
	app.mintKeeper = mint.NewKeeper(app.cdc,
		app.keyMint,
		app.paramsKeeper.Subspace(mint.DefaultParamspace),
		&stakingKeeper,
		app.feeCollectionKeeper)
	app.stakingKeeper = *stakingKeeper.SetHooks(NewStakingHooks(app.distributionKeeper.Hooks(),
		app.slashingKeeper.Hooks()))
	app.crisisKeeper = crisis.NewKeeper(app.paramsKeeper.Subspace(crisis.DefaultParamspace),
		app.distributionKeeper,
		app.bankKeeper,
		app.feeCollectionKeeper)

	app.depositKeeper = deposit.NewKeeper(app.cdc,
		app.keyDeposit,
		app.bankKeeper)
	app.vpnKeeper = vpn.NewKeeper(app.cdc,
		app.keyVPNNode,
		app.keyVPNSubscription,
		app.keyVPNSession,
		app.paramsKeeper.Subspace(vpn.DefaultParamspace),
		app.depositKeeper)

	bank.RegisterInvariants(&app.crisisKeeper,
		app.accountKeeper)
	distribution.RegisterInvariants(&app.crisisKeeper,
		app.distributionKeeper,
		app.stakingKeeper)
	_staking.RegisterInvariants(&app.crisisKeeper,
		app.stakingKeeper,
		app.feeCollectionKeeper,
		app.distributionKeeper,
		app.accountKeeper,
		app.depositKeeper)

	app.Router().
		AddRoute(bank.RouterKey, bank.NewHandler(app.bankKeeper)).
		AddRoute(staking.RouterKey, staking.NewHandler(app.stakingKeeper)).
		AddRoute(slashing.RouterKey, slashing.NewHandler(app.slashingKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewHandler(app.distributionKeeper)).
		AddRoute(gov.RouterKey, gov.NewHandler(app.govKeeper)).
		AddRoute(crisis.RouterKey, crisis.NewHandler(app.crisisKeeper)).
		AddRoute(vpn.RouterKey, vpn.NewHandler(app.vpnKeeper))

	app.QueryRouter().
		AddRoute(auth.QuerierRoute, auth.NewQuerier(app.accountKeeper)).
		AddRoute(staking.QuerierRoute, staking.NewQuerier(app.stakingKeeper, app.cdc)).
		AddRoute(slashing.QuerierRoute, slashing.NewQuerier(app.slashingKeeper, app.cdc)).
		AddRoute(distribution.QuerierRoute, distribution.NewQuerier(app.distributionKeeper)).
		AddRoute(gov.QuerierRoute, gov.NewQuerier(app.govKeeper)).
		AddRoute(mint.QuerierRoute, mint.NewQuerier(app.mintKeeper)).
		AddRoute(deposit.QuerierRoute, deposit.NewQuerier(app.depositKeeper, app.cdc)).
		AddRoute(vpn.QuerierRoute, vpn.NewQuerier(app.vpnKeeper, app.cdc))

	app.MountStores(app.keyMain, app.keyParams,
		app.keyAccount, app.keyFeeCollection,
		app.keyStaking, app.keySlashing,
		app.keyDistribution, app.keyGov, app.keyMint,
		app.keyDeposit, app.keyVPNNode, app.keyVPNSession, app.keyVPNSubscription,
		app.tkeyParams, app.tkeyStaking, app.tkeyDistribution)
	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(app.keyMain); err != nil {
			common.Exit(err.Error())
		}
	}

	return app
}

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	csdk.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	crisis.RegisterCodec(cdc)

	vpn.RegisterCodec(cdc)
	return cdc
}

func (app *Hub) BeginBlocker(ctx csdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	mint.BeginBlocker(ctx, app.mintKeeper)
	distribution.BeginBlocker(ctx, req, app.distributionKeeper)
	tags := slashing.BeginBlocker(ctx, req, app.slashingKeeper)

	return abci.ResponseBeginBlock{
		Tags: tags.ToKVPairs(),
	}
}

func (app *Hub) EndBlocker(ctx csdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	tags := gov.EndBlocker(ctx, app.govKeeper)
	validatorUpdates, endBlockerTags := staking.EndBlocker(ctx, app.stakingKeeper)
	tags = append(tags, endBlockerTags...)

	vpnTags := vpn.EndBlock(ctx, app.vpnKeeper)
	tags = tags.AppendTags(vpnTags)

	if app.invCheckPeriod != 0 && ctx.BlockHeight()%int64(app.invCheckPeriod) == 0 {
		app.assertRuntimeInvariants()
	}

	return abci.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
		Tags:             tags,
	}
}

func (app *Hub) initFromGenesisState(ctx csdk.Context, genesisState GenesisState) []abci.ValidatorUpdate {
	genesisState.Sanitize()

	for _, gacc := range genesisState.Accounts {
		acc := gacc.ToAccount()
		acc = app.accountKeeper.NewAccount(ctx, acc)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	distribution.InitGenesis(ctx, app.distributionKeeper, genesisState.Distribution)

	validators, err := staking.InitGenesis(ctx, app.stakingKeeper, genesisState.Staking)
	if err != nil {
		panic(err)
	}

	auth.InitGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper, genesisState.Auth)
	bank.InitGenesis(ctx, app.bankKeeper, genesisState.Bank)
	slashing.InitGenesis(ctx, app.slashingKeeper, genesisState.Slashing, genesisState.Staking.Validators.ToSDKValidators())
	gov.InitGenesis(ctx, app.govKeeper, genesisState.Gov)
	mint.InitGenesis(ctx, app.mintKeeper, genesisState.Mint)
	crisis.InitGenesis(ctx, app.crisisKeeper, genesisState.Crisis)

	deposit.InitGenesis(ctx, app.depositKeeper, genesisState.Deposit)
	vpn.InitGenesis(ctx, app.vpnKeeper, genesisState.VPN)

	if err = ValidateGenesisState(genesisState); err != nil {
		panic(err)
	}

	if len(genesisState.GenTxs) > 0 {
		for _, genTx := range genesisState.GenTxs {
			var tx auth.StdTx
			if err = app.cdc.UnmarshalJSON(genTx, &tx); err != nil {
				panic(err)
			}
			bz := app.cdc.MustMarshalBinaryLengthPrefixed(tx)
			res := app.BaseApp.DeliverTx(bz)
			if !res.IsOK() {
				panic(res.Log)
			}
		}

		validators = app.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	}

	return validators
}

func (app *Hub) initChainer(ctx csdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes

	var genesisState GenesisState
	if err := app.cdc.UnmarshalJSON(stateJSON, &genesisState); err != nil {
		panic(err)
	}

	validators := app.initFromGenesisState(ctx, genesisState)

	if len(req.Validators) > 0 {
		if len(req.Validators) != len(validators) {
			panic(fmt.Errorf("len(RequestInitChain.Validators) != len(validators) (%d != %d)",
				len(req.Validators), len(validators)))
		}
		sort.Sort(abci.ValidatorUpdates(req.Validators))
		sort.Sort(abci.ValidatorUpdates(validators))
		for i, val := range validators {
			if !val.Equal(req.Validators[i]) {
				panic(fmt.Errorf("validators[%d] != req.Validators[%d] ", i, i))
			}
		}
	}

	app.assertRuntimeInvariants()

	return abci.ResponseInitChain{
		Validators: validators,
	}
}

func (app *Hub) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

type StakingHooks struct {
	dh distribution.Hooks
	sh slashing.Hooks
}

func NewStakingHooks(dh distribution.Hooks, sh slashing.Hooks) StakingHooks {
	return StakingHooks{dh, sh}
}

func (h StakingHooks) AfterValidatorCreated(ctx csdk.Context, valAddr csdk.ValAddress) {
	h.dh.AfterValidatorCreated(ctx, valAddr)
	h.sh.AfterValidatorCreated(ctx, valAddr)
}

func (h StakingHooks) BeforeValidatorModified(ctx csdk.Context, valAddr csdk.ValAddress) {
	h.dh.BeforeValidatorModified(ctx, valAddr)
	h.sh.BeforeValidatorModified(ctx, valAddr)
}

func (h StakingHooks) AfterValidatorRemoved(ctx csdk.Context, consAddr csdk.ConsAddress, valAddr csdk.ValAddress) {
	h.dh.AfterValidatorRemoved(ctx, consAddr, valAddr)
	h.sh.AfterValidatorRemoved(ctx, consAddr, valAddr)
}

func (h StakingHooks) AfterValidatorBonded(ctx csdk.Context, consAddr csdk.ConsAddress, valAddr csdk.ValAddress) {
	h.dh.AfterValidatorBonded(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBonded(ctx, consAddr, valAddr)
}

// nolint: lll
func (h StakingHooks) AfterValidatorBeginUnbonding(ctx csdk.Context, consAddr csdk.ConsAddress, valAddr csdk.ValAddress) {
	h.dh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
}

func (h StakingHooks) BeforeDelegationCreated(ctx csdk.Context, delAddr csdk.AccAddress, valAddr csdk.ValAddress) {
	h.dh.BeforeDelegationCreated(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationCreated(ctx, delAddr, valAddr)
}

// nolint: lll
func (h StakingHooks) BeforeDelegationSharesModified(ctx csdk.Context, delAddr csdk.AccAddress, valAddr csdk.ValAddress) {
	h.dh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
}

func (h StakingHooks) BeforeDelegationRemoved(ctx csdk.Context, delAddr csdk.AccAddress, valAddr csdk.ValAddress) {
	h.dh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
}

func (h StakingHooks) AfterDelegationModified(ctx csdk.Context, delAddr csdk.AccAddress, valAddr csdk.ValAddress) {
	h.dh.AfterDelegationModified(ctx, delAddr, valAddr)
	h.sh.AfterDelegationModified(ctx, delAddr, valAddr)
}

func (h StakingHooks) BeforeValidatorSlashed(ctx csdk.Context, valAddr csdk.ValAddress, fraction csdk.Dec) {
	h.dh.BeforeValidatorSlashed(ctx, valAddr, fraction)
	h.sh.BeforeValidatorSlashed(ctx, valAddr, fraction)
}
