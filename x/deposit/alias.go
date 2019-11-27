// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/deposit/types/
// ALIASGEN: github.com/sentinel-official/hub/x/deposit/keeper/
// ALIASGEN: github.com/sentinel-official/hub/x/deposit/querier/
package deposit

import (
	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/querier"
	"github.com/sentinel-official/hub/x/deposit/types"
)

const (
	Codespace             = types.Codespace
	ModuleName            = types.ModuleName
	StoreKey              = types.StoreKey
	RouterKey             = types.RouterKey
	QuerierRoute          = types.QuerierRoute
	QueryDepositOfAddress = types.QueryDepositOfAddress
	QueryAllDeposits      = types.QueryAllDeposits
)

var (
	// functions aliases
	ErrorMarshal                   = types.ErrorMarshal
	ErrorUnmarshal                 = types.ErrorUnmarshal
	ErrorInvalidQueryType          = types.ErrorInvalidQueryType
	ErrorInsufficientDepositFunds  = types.ErrorInsufficientDepositFunds
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	DepositKey                     = types.DepositKey
	NewQueryDepositOfAddressParams = types.NewQueryDepositOfAddressParams
	NewKeeper                      = keeper.NewKeeper
	NewQuerier                     = querier.NewQuerier

	// variable aliases
	ModuleCdc        = types.ModuleCdc
	DepositKeyPrefix = types.DepositKeyPrefix
)

type (
	Deposit                    = types.Deposit
	GenesisState               = types.GenesisState
	QueryDepositOfAddressPrams = types.QueryDepositOfAddressPrams
	Keeper                     = keeper.Keeper
)
