package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(subscription.GetID())
	)

	value, err := k.cdc.MarshalInterface(subscription)
	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}

func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription types.Subscription, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return subscription, false
	}
	if err := k.cdc.UnmarshalInterface(value, &subscription); err != nil {
		panic(err)
	}

	return subscription, true
}

func (k *Keeper) DeleteSubscription(ctx sdk.Context, id uint64) {
	key := types.SubscriptionKey(id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetSubscriptions(ctx sdk.Context) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Subscription
		if err := k.cdc.UnmarshalInterface(iter.Value(), &item); err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForAccountKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForAccountKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for account key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForNode(ctx sdk.Context, addr hubtypes.NodeAddress) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForNodeKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForNodeKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for node key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForPlan(ctx sdk.Context, id uint64) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForPlanKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForPlanKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for plan key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForExpiryAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SubscriptionForExpiryAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteSubscriptionForExpiryAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SubscriptionForExpiryAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateSubscriptionsForExpiryAt(ctx sdk.Context, endTime time.Time, fn func(index int, item types.Subscription) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.SubscriptionForExpiryAtKeyPrefix, sdk.PrefixEndBytes(types.GetSubscriptionForExpiryAtKeyPrefix(endTime)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		subscription, found := k.GetSubscription(ctx, types.IDFromSubscriptionForExpiryAtKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for expiry at key %X does not exist", iter.Key()))
		}

		if stop := fn(i, subscription); stop {
			break
		}
		i++
	}
}

func (k *Keeper) CreateSubscriptionForNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, gigabytes, hours int64, denom string) (*types.NodeSubscription, error) {
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	var (
		count = k.GetCount(ctx)
		quota = types.Quota{
			Address:        accAddr.String(),
			AllocatedBytes: hubtypes.Gigabyte.MulRaw(gigabytes),
			ConsumedBytes:  sdk.ZeroInt(),
		}
		subscription = &types.NodeSubscription{
			BaseSubscription: &types.BaseSubscription{
				ID:       count + 1,
				Address:  accAddr.String(),
				ExpiryAt: time.Time{},
				Status:   hubtypes.StatusActive,
				StatusAt: ctx.BlockTime(),
			},
			NodeAddress: nodeAddr.String(),
			Gigabytes:   gigabytes,
			Hours:       hours,
			Deposit:     sdk.Coin{},
		}
	)

	if gigabytes != 0 {
		price, found := node.GigabytePrice(denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(denom)
		}

		subscription.ExpiryAt = ctx.BlockTime().Add(types.Year) // TODO: take the value from params?
		subscription.Deposit = sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(gigabytes),
		)
	}
	if hours != 0 {
		price, found := node.HourlyPrice(denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(denom)
		}

		subscription.ExpiryAt = ctx.BlockTime().Add(
			time.Duration(hours) * time.Hour,
		)
		subscription.Deposit = sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(hours),
		)
	}

	if err := k.AddDeposit(ctx, accAddr, subscription.Deposit); err != nil {
		return nil, err
	}

	k.SetCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.GetID())
	k.SetSubscriptionForNode(ctx, nodeAddr, subscription.GetID())
	k.SetSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())
	k.SetQuota(ctx, subscription.GetID(), quota)

	return subscription, nil
}

func (k *Keeper) CreateSubscriptionForPlan(ctx sdk.Context, accAddr sdk.AccAddress, id uint64, denom string) (*types.PlanSubscription, error) {
	plan, found := k.GetPlan(ctx, id)
	if !found {
		return nil, types.NewErrorPlanNotFound(id)
	}
	if !plan.Status.Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	price, found := plan.Price(denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(denom)
	}

	provAddr := plan.GetAddress()
	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), price); err != nil {
		return nil, err
	}

	// TODO: distribute revenue share

	var (
		count = k.GetCount(ctx)
		quota = types.Quota{
			Address:        accAddr.String(),
			AllocatedBytes: plan.Bytes,
			ConsumedBytes:  sdk.ZeroInt(),
		}
		subscription = &types.PlanSubscription{
			BaseSubscription: &types.BaseSubscription{
				ID:       count + 1,
				Address:  accAddr.String(),
				ExpiryAt: ctx.BlockTime().Add(plan.Duration),
				Status:   hubtypes.StatusActive,
				StatusAt: ctx.BlockTime(),
			},
			PlanID: plan.ID,
			Denom:  price.Denom,
		}
	)

	k.SetCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.GetID())
	k.SetSubscriptionForPlan(ctx, plan.ID, subscription.GetID())
	k.SetSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())
	k.SetQuota(ctx, subscription.GetID(), quota)

	return subscription, nil
}
