package bank

import (
	sdk "github.com/irisnet/service-sdk-go/types"
)

// expose bank module api for user
type BankI interface {
	sdk.Module
}
