package types

import (
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName = "provider"
)

var (
	TypeMsgRegisterRequest = ModuleName + ":register"
	TypeMsgUpdateRequest   = ModuleName + ":update"
)

var (
	ProviderKeyPrefix         = []byte{0x10}
	ActiveProviderKeyPrefix   = append(ProviderKeyPrefix, 0x11)
	InactiveProviderKeyPrefix = append(ProviderKeyPrefix, 0x12)
)

func ActiveProviderKey(addr hubtypes.ProvAddress) []byte {
	return append(ActiveProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveProviderKey(addr hubtypes.ProvAddress) (v []byte) {
	return append(InactiveProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}
