package token

import (
	sdk "github.com/irisnet/service-sdk-go/types"
)

type TokenI interface {
	sdk.Module

	QueryToken(symbol string) (sdk.Token, error)
}
