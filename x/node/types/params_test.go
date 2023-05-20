package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		Deposit           sdk.Coin
		ExpiryDuration    time.Duration
		GigabyteMaxPrices sdk.Coins
		GigabyteMinPrices sdk.Coins
		HourlyMaxPrices   sdk.Coins
		HourlyMinPrices   sdk.Coins
		LeaseMaxGigabytes int64
		LeaseMinGigabytes int64
		LeaseMaxHours     int64
		LeaseMinHours     int64
		RevenueShare      sdk.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"deposit empty",
			fields{
				Deposit: sdk.Coin{},
			},
			true,
		},
		{
			"deposit empty denom",
			fields{
				Deposit: sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"deposit invalid denom",
			fields{
				Deposit: sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"deposit empty amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"deposit negative amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"deposit zero amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"deposit positive amount",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"expiry_duration negative",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: -1000,
			},
			true,
		},
		{
			"expiry_duration zero",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 0,
			},
			true,
		},
		{
			"expiry_duration positive",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"gigabyte_max_prices nil",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: nil,
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"gigabyte_max_prices empty",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: sdk.Coins{},
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"gigabyte_max_prices empty denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"gigabyte_max_prices invalid denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"gigabyte_max_prices empty amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"gigabyte_max_prices negative amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"gigabyte_max_prices zero amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"gigabyte_max_prices positive amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"gigabyte_min_prices nil",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: nil,
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"gigabyte_min_prices empty",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: sdk.Coins{},
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"gigabyte_min_prices empty denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"gigabyte_min_prices invalid denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"gigabyte_min_prices empty amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"gigabyte_min_prices negative amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"gigabyte_min_prices zero amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"gigabyte_min_prices positive amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				GigabyteMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"hourly_max_prices nil",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: nil,
				RevenueShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"hourly_max_prices empty",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: sdk.Coins{},
				RevenueShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"hourly_max_prices empty denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"hourly_max_prices invalid denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"hourly_max_prices empty amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"hourly_max_prices negative amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"hourly_max_prices zero amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"hourly_max_prices positive amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMaxPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"hourly_min_prices nil",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: nil,
				RevenueShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"hourly_min_prices empty",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: sdk.Coins{},
				RevenueShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"hourly_min_prices empty denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"hourly_min_prices invalid denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"hourly_min_prices empty amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"hourly_min_prices negative amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"hourly_min_prices zero amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"hourly_min_prices positive amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:  1000,
				HourlyMinPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_max_gigabytes negative",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				LeaseMaxGigabytes: -1000,
			},
			true,
		},
		{
			"lease_max_gigabytes zero",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				LeaseMaxGigabytes: 0,
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_max_gigabytes positive",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				LeaseMaxGigabytes: 1000,
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_min_gigabytes negative",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				LeaseMinGigabytes: -1000,
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"lease_min_gigabytes zero",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				LeaseMinGigabytes: 0,
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_min_gigabytes positive",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration:    1000,
				LeaseMinGigabytes: 1000,
				RevenueShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_max_hours negative",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				LeaseMaxHours:  -1000,
			},
			true,
		},
		{
			"lease_max_hours zero",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				LeaseMaxHours:  0,
				RevenueShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_max_hours positive",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				LeaseMaxHours:  1000,
				RevenueShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_min_hours negative",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				LeaseMinHours:  -1000,
			},
			true,
		},
		{
			"lease_min_hours zero",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				LeaseMinHours:  0,
				RevenueShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"lease_min_hours positive",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				LeaseMinHours:  1000,
				RevenueShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"revenue_share empty",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.Dec{},
			},
			true,
		},
		{
			"revenue_share -10",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(-10, 0),
			},
			true,
		},
		{
			"revenue_share -1",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(-1, 0),
			},
			true,
		},
		{
			"revenue_share -0.5",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(-5, 1),
			},
			true,
		},
		{
			"revenue_share 0",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(0, 0),
			},
			false,
		},
		{
			"revenue_share 0.5",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(5, 1),
			},
			false,
		},
		{
			"revenue_share 1",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"revenue_share 10",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ExpiryDuration: 1000,
				RevenueShare:   sdk.NewDecWithPrec(10, 0),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Params{
				Deposit:           tt.fields.Deposit,
				ExpiryDuration:    tt.fields.ExpiryDuration,
				GigabyteMaxPrices: tt.fields.GigabyteMaxPrices,
				GigabyteMinPrices: tt.fields.GigabyteMinPrices,
				HourlyMaxPrices:   tt.fields.HourlyMaxPrices,
				HourlyMinPrices:   tt.fields.HourlyMinPrices,
				LeaseMaxGigabytes: tt.fields.LeaseMaxGigabytes,
				LeaseMinGigabytes: tt.fields.LeaseMinGigabytes,
				LeaseMaxHours:     tt.fields.LeaseMaxHours,
				LeaseMinHours:     tt.fields.LeaseMinHours,
				RevenueShare:      tt.fields.RevenueShare,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
