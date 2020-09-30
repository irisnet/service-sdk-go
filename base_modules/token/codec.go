package token

import (
	"github.com/irisnet/service-sdk-go/codec"
	"github.com/irisnet/service-sdk-go/codec/types"
	cryptocodec "github.com/irisnet/service-sdk-go/crypto/codec"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface("irismod.token.TokenI", (*TokenInterface)(nil), &Token{})
}
