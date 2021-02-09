package types

import (
	cdctypes "github.com/irisnet/service-sdk-go/codec/types"
	cryptotypes "github.com/irisnet/service-sdk-go/crypto/types"
)

type Response interface {
	Convert() interface{}
}

type SplitAble interface {
	Len() int
	Sub(begin, end int) SplitAble
}

type Module interface {
	Name() string
	RegisterInterfaceTypes(registry cdctypes.InterfaceRegistry)
}

type KeyManager interface {
	Sign(name, password string, data []byte) ([]byte, cryptotypes.PubKey, error)
	Insert(name, password string) (string, string, error)
	Recover(name, password, mnemonic string) (string, error)
	Import(name, password string, privKeyArmor string) (address string, err error)
	Export(name, password string) (privKeyArmor string, err error)
	Delete(name, password string) error
	Find(name, password string) (cryptotypes.PubKey, AccAddress, error)
}
