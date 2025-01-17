package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgCancelRequest)(nil)
	_ sdk.Msg = (*MsgAllocateRequest)(nil)
)

func NewMsgCancelRequest(from sdk.AccAddress, id uint64) *MsgCancelRequest {
	return &MsgCancelRequest{
		From: from.String(),
		ID:   id,
	}
}

func (m *MsgCancelRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgCancelRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgAllocateRequest(from sdk.AccAddress, id uint64, addr sdk.AccAddress, bytes sdk.Int) *MsgAllocateRequest {
	return &MsgAllocateRequest{
		From:    from.String(),
		ID:      id,
		Address: addr.String(),
		Bytes:   bytes,
	}
}

func (m *MsgAllocateRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidMessage, "address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidMessage, "invalid address %s", err)
	}
	if m.Bytes.IsNil() {
		return errors.Wrap(ErrorInvalidMessage, "bytes cannot be nil")
	}
	if m.Bytes.IsNegative() {
		return errors.Wrap(ErrorInvalidMessage, "bytes cannot be negative")
	}

	return nil
}

func (m *MsgAllocateRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
