package node

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	var (
		gigabyteMaxPricesModified = k.IsGigabyteMaxPricesModified(ctx)
		gigabyteMinPricesModified = k.IsGigabyteMinPricesModified(ctx)
		hourlyMaxPricesModified   = k.IsHourlyMaxPricesModified(ctx)
		hourlyMinPricesModified   = k.IsHourlyMinPricesModified(ctx)
	)

	if gigabyteMaxPricesModified || gigabyteMinPricesModified || hourlyMaxPricesModified || hourlyMinPricesModified {
		gigabyteMaxPrices := sdk.NewCoins()
		if gigabyteMaxPricesModified {
			gigabyteMaxPrices = k.GigabyteMaxPrices(ctx)
		}

		gigabyteMinPrices := sdk.NewCoins()
		if gigabyteMinPricesModified {
			gigabyteMinPrices = k.GigabyteMinPrices(ctx)
		}

		hourlyMaxPrices := sdk.NewCoins()
		if hourlyMaxPricesModified {
			hourlyMaxPrices = k.HourlyMaxPrices(ctx)
		}

		hourlyMinPrices := sdk.NewCoins()
		if hourlyMinPricesModified {
			hourlyMinPrices = k.HourlyMinPrices(ctx)
		}

		k.IterateNodes(ctx, func(_ int, item types.Node) bool {
			if item.GigabytePrices != nil {
				if gigabyteMaxPricesModified {
					for _, coin := range gigabyteMaxPrices {
						amount := item.GigabytePrices.AmountOf(coin.Denom)
						if amount.GT(coin.Amount) {
							item.GigabytePrices = item.GigabytePrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}

				if gigabyteMinPricesModified {
					for _, coin := range gigabyteMinPrices {
						amount := item.GigabytePrices.AmountOf(coin.Denom)
						if amount.LT(coin.Amount) {
							item.GigabytePrices = item.GigabytePrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}
			}

			if item.HourlyPrices != nil {
				if hourlyMaxPricesModified {
					for _, coin := range hourlyMaxPrices {
						amount := item.HourlyPrices.AmountOf(coin.Denom)
						if amount.GT(coin.Amount) {
							item.HourlyPrices = item.HourlyPrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}

				if hourlyMinPricesModified {
					for _, coin := range hourlyMinPrices {
						amount := item.HourlyPrices.AmountOf(coin.Denom)
						if amount.LT(coin.Amount) {
							item.HourlyPrices = item.HourlyPrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}
			}

			k.SetNode(ctx, item)
			return false
		})
	}

	k.IterateNodesForExpiryAt(ctx, ctx.BlockTime(), func(_ int, item types.Node) bool {
		nodeAddr := item.GetAddress()
		k.DeleteActiveNode(ctx, nodeAddr)
		k.DeleteNodeForExpiryAt(ctx, item.ExpiryAt, nodeAddr)

		item.ExpiryAt = time.Time{}
		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()

		k.SetNode(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				Address: item.Address,
				Status:  item.Status,
			},
		)

		return false
	})

	return nil
}
