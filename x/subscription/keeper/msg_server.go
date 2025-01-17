package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}
	if !fromAddr.Equals(subscription.GetAddress()) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())

	expiryDuration := k.ExpiryDuration(ctx)
	subscription.SetExpiryAt(
		subscription.GetExpiryAt().Add(expiryDuration),
	)
	subscription.SetStatus(hubtypes.StatusInactivePending)
	subscription.SetStatusAt(ctx.BlockTime())

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			ID:     subscription.GetID(),
			Status: subscription.GetStatus(),
		},
	)

	payout, found := k.GetPayout(ctx, subscription.GetID())
	if !found {
		return &types.MsgCancelResponse{}, nil
	}

	k.DeletePayoutForTimestamp(ctx, payout.Timestamp, payout.ID)

	payout.Timestamp = time.Time{}
	k.SetPayout(ctx, payout)

	return &types.MsgCancelResponse{}, nil
}

func (k *msgServer) MsgAllocate(c context.Context, msg *types.MsgAllocateRequest) (*types.MsgAllocateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if subscription.Type() != types.TypePlan {
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type())
	}
	if !fromAddr.Equals(subscription.GetAddress()) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	fromAlloc, found := k.GetAllocation(ctx, subscription.GetID(), fromAddr)
	if !found {
		return nil, types.NewErrorAllocationNotFound(subscription.GetID(), fromAddr)
	}

	toAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	toAlloc, found := k.GetAllocation(ctx, subscription.GetID(), toAddr)
	if !found {
		toAlloc = types.Allocation{
			ID:            subscription.GetID(),
			Address:       toAddr.String(),
			GrantedBytes:  sdk.ZeroInt(),
			UtilisedBytes: sdk.ZeroInt(),
		}

		k.SetSubscriptionForAccount(ctx, toAddr, subscription.GetID())
	}

	var (
		granted   = fromAlloc.GrantedBytes.Add(toAlloc.GrantedBytes)
		utilised  = fromAlloc.UtilisedBytes.Add(toAlloc.UtilisedBytes)
		available = granted.Sub(utilised)
	)

	if msg.Bytes.GT(available) {
		return nil, types.NewErrorInsufficientBytes(subscription.GetID(), msg.Bytes)
	}

	fromAlloc.GrantedBytes = available.Sub(msg.Bytes)
	if fromAlloc.GrantedBytes.LT(fromAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), fromAddr)
	}

	k.SetAllocation(ctx, fromAlloc)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: fromAlloc.Address,
			Bytes:   fromAlloc.GrantedBytes,
		},
	)

	toAlloc.GrantedBytes = msg.Bytes
	if toAlloc.GrantedBytes.LT(toAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), toAddr)
	}

	k.SetAllocation(ctx, toAlloc)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: toAlloc.Address,
			Bytes:   toAlloc.GrantedBytes,
		},
	)

	return &types.MsgAllocateResponse{}, nil
}
