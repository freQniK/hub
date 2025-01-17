// DO NOT COVER

package simulation

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	nodesimulation "github.com/sentinel-official/hub/x/node/simulation"
	plansimulation "github.com/sentinel-official/hub/x/plan/simulation"
	providersimulation "github.com/sentinel-official/hub/x/provider/simulation"
	sessionsimulation "github.com/sentinel-official/hub/x/session/simulation"
	subscriptionsimulation "github.com/sentinel-official/hub/x/subscription/simulation"
	"github.com/sentinel-official/hub/x/vpn/expected"
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func WeightedOperations(
	cdc codec.Codec,
	txConfig client.TxConfig,
	params simtypes.AppParams,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) []simtypes.WeightedOperation {
	var operations []simtypes.WeightedOperation
	operations = append(operations, providersimulation.WeightedOperations(cdc, txConfig, params, ak, bk, k.Provider)...)
	operations = append(operations, nodesimulation.WeightedOperations(cdc, txConfig, params, ak, bk, k.Node)...)
	operations = append(operations, plansimulation.WeightedOperations(cdc, txConfig, params, ak, bk, k.Plan)...)
	operations = append(operations, subscriptionsimulation.WeightedOperations(cdc, txConfig, params, ak, bk, k.Subscription)...)
	operations = append(operations, sessionsimulation.WeightedOperations(cdc, txConfig, params, ak, bk, k.Session)...)

	return operations
}
