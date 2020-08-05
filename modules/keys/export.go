package keys

import (
	sdk "github.com/irisnet/service-sdk-go/types"
)

type KeyI interface {
	sdk.Module
	Add(name, password string) (address string, mnemonic string, err sdk.Error)
	Recover(name, password, mnemonic string) (address string, err sdk.Error)
	Import(name, password, privKeyArmor string) (address string, err sdk.Error)
	Export(name, password string) (privKeyArmor string, err sdk.Error)
	Delete(name, password string) sdk.Error
	Show(name, password string) (string, sdk.Error)
}
