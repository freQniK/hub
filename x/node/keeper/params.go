package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) Deposit(ctx sdk.Context) (v sdk.Coin) {
	k.params.Get(ctx, types.KeyDeposit, &v)
	return
}

func (k *Keeper) ExpiryDuration(ctx sdk.Context) (v time.Duration) {
	k.params.Get(ctx, types.KeyExpiryDuration, &v)
	return
}

func (k *Keeper) GigabyteMaxPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyGigabyteMaxPrices, &v)
	return
}

func (k *Keeper) GigabyteMinPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyGigabyteMinPrices, &v)
	return
}

func (k *Keeper) HourlyMaxPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyHourlyMaxPrices, &v)
	return
}

func (k *Keeper) HourlyMinPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyHourlyMinPrices, &v)
	return
}

func (k *Keeper) LeaseMaxGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyLeaseMaxGigabytes, &v)
	return
}

func (k *Keeper) LeaseMinGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyLeaseMinGigabytes, &v)
	return
}

func (k *Keeper) LeaseMaxHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyLeaseMaxHours, &v)
	return
}

func (k *Keeper) LeaseMinHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyLeaseMinHours, &v)
	return
}

func (k *Keeper) LeaseDistributionDuration(ctx sdk.Context) (v time.Duration) {
	k.params.Get(ctx, types.KeyLeaseDistributionDuration, &v)
	return
}

func (k *Keeper) RevenueShare(ctx sdk.Context) (v sdk.Dec) {
	k.params.Get(ctx, types.KeyRevenueShare, &v)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Deposit(ctx),
		k.ExpiryDuration(ctx),
		k.GigabyteMaxPrices(ctx),
		k.GigabyteMinPrices(ctx),
		k.HourlyMaxPrices(ctx),
		k.HourlyMinPrices(ctx),
		k.LeaseMaxGigabytes(ctx),
		k.LeaseMinGigabytes(ctx),
		k.LeaseMaxHours(ctx),
		k.LeaseMinHours(ctx),
		k.LeaseDistributionDuration(ctx),
		k.RevenueShare(ctx),
	)
}

func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices sdk.Coins) bool {
	maxPrices := k.GigabyteMaxPrices(ctx)
	for _, coin := range maxPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrices := k.GigabyteMinPrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidHourlyPrices(ctx sdk.Context, prices sdk.Coins) bool {
	maxPrices := k.HourlyMaxPrices(ctx)
	for _, coin := range maxPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrices := k.HourlyMinPrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidLeaseGigabytes(ctx sdk.Context, gigabytes int64) bool {
	maxGigabytes := k.LeaseMaxGigabytes(ctx)
	if maxGigabytes != 0 && gigabytes > maxGigabytes {
		return false
	}

	minGigabytes := k.LeaseMinGigabytes(ctx)
	if minGigabytes != 0 && gigabytes < minGigabytes {
		return false
	}

	return true
}

func (k *Keeper) IsValidLeaseHours(ctx sdk.Context, hours int64) bool {
	maxHours := k.LeaseMaxHours(ctx)
	if maxHours != 0 && hours > maxHours {
		return false
	}

	minHours := k.LeaseMinHours(ctx)
	if minHours != 0 && hours < minHours {
		return false
	}

	return true
}

func (k *Keeper) IsGigabyteMaxPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyGigabyteMaxPrices)
}

func (k *Keeper) IsGigabyteMinPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyGigabyteMinPrices)
}

func (k *Keeper) IsHourlyMaxPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyHourlyMaxPrices)
}

func (k *Keeper) IsHourlyMinPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyHourlyMinPrices)
}
