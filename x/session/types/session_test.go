package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestSession_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty",
			fields{
				Address: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Address: hubtypes.TestBech32AccAddr20Bytes,
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				Address: tt.fields.Address,
			}
			if got := m.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_GetNodeAddress(t *testing.T) {
	type fields struct {
		NodeAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
	}{
		{
			"empty",
			fields{
				NodeAddress: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				NodeAddress: tt.fields.NodeAddress,
			}
			if got := m.GetNodeAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNodeAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Validate(t *testing.T) {
	type fields struct {
		ID             uint64
		SubscriptionID uint64
		NodeAddress    string
		Address        string
		Bandwidth      hubtypes.Bandwidth
		Duration       time.Duration
		ExpiryAt       time.Time
		Status         hubtypes.Status
		StatusAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"id zero",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"id positive",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"subscription_id zero",
			fields{
				ID:             1000,
				SubscriptionID: 0,
			},
			true,
		},
		{
			"subscription_id positive",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"node_address empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "",
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "invalid",
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr10Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr30Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"address empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"address 20 bytes",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"upload empty and download empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.Int{}, Download: sdk.Int{}},
			},
			true,
		},
		{
			"upload empty and download non-empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.Int{}, Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"upload non-empty and download empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.Int{}},
			},
			true,
		},
		{
			"upload non-empty and download non-empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"upload negative and download negative",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"upload negative and download zero",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative upload and download positive",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"upload zero and download negative",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"upload zero and download zero",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"upload zero and download positive",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"upload positive and download negative",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"upload positive and download zero",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"upload positive and download positive",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"duration negative",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       -1000,
			},
			true,
		},
		{
			"duration zero",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       0,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"duration positive",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"expiry_at empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Time{},
			},
			true,
		},
		{
			"expiry_at non-empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"status unspecified",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusInactivePending,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"status inactive",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"status_at empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Time{},
			},
			true,
		},
		{
			"status_at non-empty",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    hubtypes.TestBech32NodeAddr20Bytes,
				Address:        hubtypes.TestBech32AccAddr20Bytes,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Duration:       1000,
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				ID:             tt.fields.ID,
				SubscriptionID: tt.fields.SubscriptionID,
				NodeAddress:    tt.fields.NodeAddress,
				Address:        tt.fields.Address,
				Bandwidth:      tt.fields.Bandwidth,
				Duration:       tt.fields.Duration,
				ExpiryAt:       tt.fields.ExpiryAt,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
