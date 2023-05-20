package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultDeposit                 = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))
	DefaultExpiryDuration          = 30 * time.Second
	DefaultGigabyteMaxPrices       = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))
	DefaultGigabyteMinPrices       = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)))
	DefaultHourlyMaxPrices         = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))
	DefaultHourlyMinPrices         = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)))
	DefaultLeaseMaxGigabytes int64 = 10
	DefaultLeaseMinGigabytes int64 = 1
	DefaultLeaseMaxHours     int64 = 10
	DefaultLeaseMinHours     int64 = 1
	DefaultRevenueShare            = sdk.NewDecWithPrec(1, 1)
)

var (
	KeyDeposit           = []byte("Deposit")
	KeyExpiryDuration    = []byte("ExpiryDuration")
	KeyGigabyteMaxPrices = []byte("GigabyteMaxPrices")
	KeyGigabyteMinPrices = []byte("GigabyteMinPrices")
	KeyHourlyMaxPrices   = []byte("HourlyMaxPrices")
	KeyHourlyMinPrices   = []byte("HourlyMinPrices")
	KeyLeaseMaxGigabytes = []byte("LeaseMaxGigabytes")
	KeyLeaseMinGigabytes = []byte("LeaseMinGigabytes")
	KeyLeaseMaxHours     = []byte("LeaseMaxHours")
	KeyLeaseMinHours     = []byte("LeaseMinHours")
	KeyRevenueShare      = []byte("RevenueShare")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}
	if err := validateExpiryDuration(m.ExpiryDuration); err != nil {
		return err
	}
	if err := validateGigabyteMaxPrices(m.GigabyteMaxPrices); err != nil {
		return err
	}
	if err := validateGigabyteMinPrices(m.GigabyteMinPrices); err != nil {
		return err
	}
	if err := validateHourlyMaxPrices(m.HourlyMaxPrices); err != nil {
		return err
	}
	if err := validateHourlyMinPrices(m.HourlyMinPrices); err != nil {
		return err
	}
	if err := validateLeaseMaxGigabytes(m.LeaseMaxGigabytes); err != nil {
		return err
	}
	if err := validateLeaseMinGigabytes(m.LeaseMinGigabytes); err != nil {
		return err
	}
	if err := validateLeaseMaxHours(m.LeaseMaxHours); err != nil {
		return err
	}
	if err := validateLeaseMinHours(m.LeaseMinHours); err != nil {
		return err
	}
	if err := validateRevenueShare(m.RevenueShare); err != nil {
		return err
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:         KeyDeposit,
			Value:       &m.Deposit,
			ValidatorFn: validateDeposit,
		},
		{
			Key:         KeyExpiryDuration,
			Value:       &m.ExpiryDuration,
			ValidatorFn: validateExpiryDuration,
		},
		{
			Key:         KeyGigabyteMaxPrices,
			Value:       &m.GigabyteMaxPrices,
			ValidatorFn: validateGigabyteMaxPrices,
		},
		{
			Key:         KeyGigabyteMinPrices,
			Value:       &m.GigabyteMinPrices,
			ValidatorFn: validateGigabyteMinPrices,
		},
		{
			Key:         KeyHourlyMaxPrices,
			Value:       &m.HourlyMaxPrices,
			ValidatorFn: validateHourlyMaxPrices,
		},
		{
			Key:         KeyHourlyMinPrices,
			Value:       &m.HourlyMinPrices,
			ValidatorFn: validateHourlyMinPrices,
		},
		{
			Key:         KeyLeaseMaxGigabytes,
			Value:       &m.LeaseMaxGigabytes,
			ValidatorFn: validateLeaseMaxGigabytes,
		},
		{
			Key:         KeyLeaseMinGigabytes,
			Value:       &m.LeaseMinGigabytes,
			ValidatorFn: validateLeaseMinGigabytes,
		},
		{
			Key:         KeyLeaseMaxHours,
			Value:       &m.LeaseMaxHours,
			ValidatorFn: validateLeaseMaxHours,
		},
		{
			Key:         KeyLeaseMinHours,
			Value:       &m.LeaseMinHours,
			ValidatorFn: validateLeaseMinHours,
		},
		{
			Key:         KeyRevenueShare,
			Value:       &m.RevenueShare,
			ValidatorFn: validateRevenueShare,
		},
	}
}

func NewParams(
	deposit sdk.Coin, expiryDuration time.Duration, gigabyteMaxPrices, gigabyteMinPrices,
	hourlyMaxPrices, hourlyMinPrices sdk.Coins, leaseMaxGigabytes, leaseMinGigabytes,
	leaseMaxHours, leaseMinHours int64, revenueShare sdk.Dec,
) Params {
	return Params{
		Deposit:           deposit,
		ExpiryDuration:    expiryDuration,
		GigabyteMaxPrices: gigabyteMaxPrices,
		GigabyteMinPrices: gigabyteMinPrices,
		HourlyMaxPrices:   hourlyMaxPrices,
		HourlyMinPrices:   hourlyMinPrices,
		LeaseMaxGigabytes: leaseMaxGigabytes,
		LeaseMinGigabytes: leaseMinGigabytes,
		LeaseMaxHours:     leaseMaxHours,
		LeaseMinHours:     leaseMinHours,
		RevenueShare:      revenueShare,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultDeposit,
		DefaultExpiryDuration,
		DefaultGigabyteMaxPrices,
		DefaultGigabyteMinPrices,
		DefaultHourlyMaxPrices,
		DefaultHourlyMinPrices,
		DefaultLeaseMaxGigabytes,
		DefaultLeaseMinGigabytes,
		DefaultLeaseMaxHours,
		DefaultLeaseMinHours,
		DefaultRevenueShare,
	)
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func validateDeposit(v interface{}) error {
	value, ok := v.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("deposit cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !value.IsValid() {
		return fmt.Errorf("invalid deposit %s", value)
	}

	return nil
}

func validateExpiryDuration(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("expiry_duration cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("expiry_duration cannot be zero")
	}

	return nil
}

func validateGigabyteMaxPrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("gigabyte_max_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("gigabyte_max_prices must be valid")
	}

	return nil
}

func validateGigabyteMinPrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("gigabyte_min_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("gigabyte_min_prices must be valid")
	}

	return nil
}

func validateHourlyMaxPrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("hourly_max_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("hourly_max_prices must be valid")
	}

	return nil
}

func validateHourlyMinPrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("hourly_min_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("hourly_min_prices must be valid")
	}

	return nil
}

func validateLeaseMaxGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("lease_max_gigabytes cannot be negative")
	}

	return nil
}

func validateLeaseMinGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("lease_min_gigabytes cannot be negative")
	}

	return nil
}

func validateLeaseMaxHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("lease_max_hours cannot be negative")
	}

	return nil
}

func validateLeaseMinHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("lease_min_hours cannot be negative")
	}

	return nil
}

func validateRevenueShare(v interface{}) error {
	value, ok := v.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("revenue_share cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("revenue_share cannot be negative")
	}
	if value.GT(sdk.OneDec()) {
		return fmt.Errorf("revenue_share cannot be greater than 1")
	}

	return nil
}
