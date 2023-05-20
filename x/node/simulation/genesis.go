// DO NOT COVER

package simulation

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/x/node/types"
)

func RandomizedGenesisState(state *module.SimulationState) *types.GenesisState {
	var (
		deposit                   sdk.Coin
		expiryDuration            time.Duration
		gigabyteMaxPrices         sdk.Coins
		gigabyteMinPrices         sdk.Coins
		hourlyMaxPrices           sdk.Coins
		hourlyMinPrices           sdk.Coins
		leaseMaxGigabytes         int64
		leaseMinGigabytes         int64
		leaseMaxHours             int64
		leaseMinHours             int64
		leaseDistributionDuration time.Duration
		revenueShare              sdk.Dec
	)

	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyDeposit),
		&deposit,
		state.Rand,
		func(r *rand.Rand) {
			deposit = sdk.NewInt64Coin(
				sdk.DefaultBondDenom,
				r.Int63n(MaxInt),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyExpiryDuration),
		&expiryDuration,
		state.Rand,
		func(r *rand.Rand) {
			expiryDuration = time.Duration(r.Int63n(MaxInt)) * time.Millisecond
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyGigabyteMaxPrices),
		&gigabyteMaxPrices,
		state.Rand,
		func(r *rand.Rand) {
			gigabyteMaxPrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyGigabyteMinPrices),
		&gigabyteMinPrices,
		state.Rand,
		func(r *rand.Rand) {
			gigabyteMinPrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyHourlyMaxPrices),
		&hourlyMaxPrices,
		state.Rand,
		func(r *rand.Rand) {
			hourlyMaxPrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyHourlyMinPrices),
		&hourlyMinPrices,
		state.Rand,
		func(r *rand.Rand) {
			hourlyMinPrices = sdk.NewCoins(
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxInt),
				),
			)
		},
	)
	state.AppParams.GetOrGenerate(
		state.Cdc,
		string(types.KeyRevenueShare),
		&revenueShare,
		state.Rand,
		func(r *rand.Rand) {
			revenueShare = sdk.NewDecWithPrec(
				r.Int63n(MaxInt),
				6,
			)
		},
	)

	return types.NewGenesisState(
		RandomNodes(state.Rand, state.Accounts),
		types.NewParams(
			deposit, expiryDuration, gigabyteMaxPrices, gigabyteMinPrices,
			hourlyMaxPrices, hourlyMinPrices, leaseMaxGigabytes, leaseMinGigabytes,
			leaseMaxHours, leaseMinHours, leaseDistributionDuration, revenueShare,
		),
	)
}
