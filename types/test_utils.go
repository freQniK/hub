package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	TestTimeZero = time.Time{}
	TestTimeNow  = time.Now()

	TestAddrEmpty         = ""
	TestAddrInvalid       = "invalid"
	TestAddrInvalidPrefix = "invalid1qypqxpq9qcrsszgszyfpx9q4zct3sxfqe52gp4"

	TestBech32AccAddr10Bytes = "sent1qypqxpq9qcrsszgslawd5s"
	TestBech32AccAddr20Bytes = "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj"
	TestBech32AccAddr30Bytes = "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8"

	TestBech32NodeAddr10Bytes = "sentnode1qypqxpq9qcrsszgse4wwrm"
	TestBech32NodeAddr20Bytes = "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey"
	TestBech32NodeAddr30Bytes = "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv"

	TestBech32ProvAddr10Bytes = "sentprov1qypqxpq9qcrsszgsutj8xr"
	TestBech32ProvAddr20Bytes = "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82"
	TestBech32ProvAddr30Bytes = "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx"

	TestDenomEmpty   = ""
	TestDenomInvalid = "i"
	TestDenomOne     = "one"

	TestIntEmpty    = sdk.Int{}
	TestIntNegative = sdk.NewInt(-1000)
	TestIntZero     = sdk.NewInt(0)
	TestIntPositive = sdk.NewInt(1000)

	TestCoinEmpty          = sdk.Coin{}
	TestCoinEmptyDenom     = sdk.Coin{Denom: TestDenomEmpty, Amount: TestIntPositive}
	TestCoinInvalidDenom   = sdk.Coin{Denom: TestDenomInvalid, Amount: TestIntPositive}
	TestCoinEmptyAmount    = sdk.Coin{Denom: TestDenomOne, Amount: TestIntEmpty}
	TestCoinNegativeAmount = sdk.Coin{Denom: TestDenomOne, Amount: TestIntNegative}
	TestCoinZeroAmount     = sdk.Coin{Denom: TestDenomOne, Amount: TestIntZero}
	TestCoinPositiveAmount = sdk.Coin{Denom: TestDenomOne, Amount: TestIntPositive}

	TestCoinsNil            sdk.Coins = nil
	TestCoinsEmpty                    = sdk.Coins{}
	TestCoinsEmptyDenom               = sdk.Coins{sdk.Coin{Denom: TestDenomEmpty, Amount: TestIntPositive}}
	TestCoinsInvalidDenom             = sdk.Coins{sdk.Coin{Denom: TestDenomInvalid, Amount: TestIntPositive}}
	TestCoinsEmptyAmount              = sdk.Coins{sdk.Coin{Denom: TestDenomOne, Amount: TestIntEmpty}}
	TestCoinsNegativeAmount           = sdk.Coins{sdk.Coin{Denom: TestDenomOne, Amount: TestIntNegative}}
	TestCoinsZeroAmount               = sdk.Coins{sdk.Coin{Denom: TestDenomOne, Amount: TestIntZero}}
	TestCoinsPositiveAmount           = sdk.Coins{sdk.Coin{Denom: TestDenomOne, Amount: TestIntPositive}}
)
