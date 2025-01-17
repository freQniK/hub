package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
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

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provAddr := hubtypes.ProvAddress(accAddr.Bytes())
	if k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorDuplicateProvider(provAddr)
	}

	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	provider := types.Provider{
		Address:     provAddr.String(),
		Name:        msg.Name,
		Identity:    msg.Identity,
		Website:     msg.Website,
		Description: msg.Description,
		Status:      hubtypes.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegister{
			Address: provider.Address,
		},
	)

	return &types.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	provAddr, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	if len(msg.Name) > 0 {
		provider.Name = msg.Name
	}

	provider.Identity = msg.Identity
	provider.Website = msg.Website
	provider.Description = msg.Description

	if !msg.Status.Equal(hubtypes.StatusUnspecified) {
		if provider.Status.Equal(hubtypes.StatusActive) {
			if msg.Status.Equal(hubtypes.StatusInactive) {
				k.DeleteActiveProvider(ctx, provAddr)
			}
		}
		if provider.Status.Equal(hubtypes.StatusInactive) {
			if msg.Status.Equal(hubtypes.StatusActive) {
				k.DeleteInactiveProvider(ctx, provAddr)
			}
		}

		provider.Status = msg.Status
		provider.StatusAt = ctx.BlockTime()
	}

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdate{
			Address: provider.Address,
		},
	)

	return &types.MsgUpdateResponse{}, nil
}
