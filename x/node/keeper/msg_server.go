package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
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

func (k *msgServer) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
		}
	}
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
		}
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr := hubtypes.NodeAddress(accAddr.Bytes())
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	node := types.Node{
		Address:        nodeAddr.String(),
		GigabytePrices: msg.GigabytePrices,
		HourlyPrices:   msg.HourlyPrices,
		RemoteURL:      msg.RemoteURL,
		ExpiryAt:       time.Time{},
		Status:         hubtypes.StatusInactive,
		StatusAt:       ctx.BlockTime(),
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegister{
			Address: node.Address,
		},
	)

	return &types.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdateDetails(c context.Context, msg *types.MsgUpdateDetailsRequest) (*types.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
		}
	}
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
		}
	}

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	node.GigabytePrices = msg.GigabytePrices
	node.HourlyPrices = msg.HourlyPrices

	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDetails{
			Address: node.Address,
		},
	)

	return &types.MsgUpdateDetailsResponse{}, nil
}

func (k *msgServer) MsgUpdateStatus(c context.Context, msg *types.MsgUpdateStatusRequest) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if node.Status.Equal(hubtypes.StatusActive) {
		k.DeleteNodeForExpiryAt(ctx, node.ExpiryAt, nodeAddr)
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddr)
		}
	}
	if node.Status.Equal(hubtypes.StatusInactive) {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
		}
	}

	if msg.Status.Equal(hubtypes.StatusActive) {
		node.ExpiryAt = ctx.BlockTime().Add(
			k.ExpiryDuration(ctx),
		)
		k.SetNodeForExpiryAt(ctx, node.ExpiryAt, nodeAddr)
	}
	if msg.Status.Equal(hubtypes.StatusInactive) {
		node.ExpiryAt = time.Time{}
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			Address: node.Address,
			Status:  node.Status,
		},
	)

	return &types.MsgUpdateStatusResponse{}, nil
}

func (k *msgServer) MsgSubscribe(c context.Context, msg *types.MsgSubscribeRequest) (*types.MsgSubscribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if msg.Gigabytes != 0 {
		if !k.IsValidLeaseGigabytes(ctx, msg.Gigabytes) {
			return nil, types.NewErrorInvalidLeaseGigabytes(msg.Gigabytes)
		}
	}
	if msg.Hours != 0 {
		if !k.IsValidLeaseHours(ctx, msg.Hours) {
			return nil, types.NewErrorInvalidLeaseHours(msg.Hours)
		}
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	subscription, err := k.CreateSubscriptionForNode(ctx, accAddr, nodeAddr, msg.Gigabytes, msg.Hours, msg.Denom)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventCreateSubscription{
			ID:          subscription.ID,
			NodeAddress: subscription.NodeAddress,
		},
	)

	lease := types.Lease{
		ID:             subscription.ID,
		Bytes:          hubtypes.Gigabyte.MulRaw(msg.Gigabytes),
		Duration:       msg.Hours * time.Hour.Nanoseconds(),
		Price:          sdk.Coin{},
		DistributionAt: time.Time{},
	}

	if msg.Gigabytes != 0 {
		lease.Price = sdk.NewCoin(
			subscription.Deposit.Denom,
			subscription.Deposit.Amount.QuoRaw(msg.Gigabytes),
		)
	}
	if msg.Hours != 0 {
		lease.Price = sdk.NewCoin(
			subscription.Deposit.Denom,
			subscription.Deposit.Amount.QuoRaw(msg.Hours),
		)
		lease.DistributionAt = ctx.BlockTime().Add(time.Hour) // TODO: take the value from params?
	}

	k.SetLease(ctx, lease)
	k.SetLeaseForAccount(ctx, accAddr, lease.ID)
	k.SetLeaseForNode(ctx, nodeAddr, lease.ID)
	if !lease.DistributionAt.IsZero() {
		k.SetLeaseForDistributionAt(ctx, lease.DistributionAt, lease.ID)
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventLease{
			ID:     lease.ID,
			Lessor: nodeAddr.String(),
			Lessee: accAddr.String(),
		},
	)

	return &types.MsgSubscribeResponse{}, nil
}
