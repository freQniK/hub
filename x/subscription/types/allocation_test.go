package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestAllocation_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty account address",
			fields{
				Address: "",
			},
			nil,
		},
		{
			"20 bytes account address",
			fields{
				Address: hubtypes.TestBech32AccAddr20Bytes,
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Allocation{
				Address: tt.fields.Address,
			}
			if got := m.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocation_Validate(t *testing.T) {
	type fields struct {
		Address       string
		GrantedBytes  sdk.Int
		UtilisedBytes sdk.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty account address",
			fields{
				Address: "",
			},
			true,
		},
		{
			"invalid account address",
			fields{
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix account address",
			fields{
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes account address",
			fields{
				Address:       hubtypes.TestBech32AccAddr10Bytes,
				GrantedBytes:  sdk.NewInt(0),
				UtilisedBytes: sdk.NewInt(0),
			},
			false,
		},
		{
			"20 bytes account address",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(0),
				UtilisedBytes: sdk.NewInt(0),
			},
			false,
		},
		{
			"30 bytes account address",
			fields{
				Address:       hubtypes.TestBech32AccAddr30Bytes,
				GrantedBytes:  sdk.NewInt(0),
				UtilisedBytes: sdk.NewInt(0),
			},
			false,
		},
		{
			"nil granted",
			fields{
				Address:      hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes: sdk.Int{},
			},
			true,
		},
		{
			"negative granted",
			fields{
				Address:      hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero granted",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(0),
				UtilisedBytes: sdk.NewInt(0),
			},
			false,
		},
		{
			"positive granted",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(1000),
				UtilisedBytes: sdk.NewInt(0),
			},
			false,
		},
		{
			"nil utilised",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(1000),
				UtilisedBytes: sdk.Int{},
			},
			true,
		},
		{
			"negative utilised",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(1000),
				UtilisedBytes: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero utilised",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(1000),
				UtilisedBytes: sdk.NewInt(0),
			},
			false,
		},
		{
			"positive utilised",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(1000),
				UtilisedBytes: sdk.NewInt(1000),
			},
			false,
		},
		{
			"granted less than utilised",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(1000),
				UtilisedBytes: sdk.NewInt(2000),
			},
			true,
		},
		{
			"granted equals to utilised",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(2000),
				UtilisedBytes: sdk.NewInt(2000),
			},
			false,
		},
		{
			"granted greater than utilised",
			fields{
				Address:       hubtypes.TestBech32AccAddr20Bytes,
				GrantedBytes:  sdk.NewInt(2000),
				UtilisedBytes: sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Allocation{
				Address:       tt.fields.Address,
				GrantedBytes:  tt.fields.GrantedBytes,
				UtilisedBytes: tt.fields.UtilisedBytes,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
