package store

import (
	"github.com/irisnet/service-sdk-go/codec"
	cryptoAmino "github.com/irisnet/service-sdk-go/crypto/codec"
	"github.com/irisnet/service-sdk-go/crypto/hd"
	cryptotypes "github.com/irisnet/service-sdk-go/crypto/types"
)

var cdc *codec.LegacyAmino

func init() {
	cdc = codec.NewLegacyAmino()
	cryptoAmino.RegisterCrypto(cdc)
	RegisterCodec(cdc)
	cdc.Seal()
}

// RegisterCodec registers concrete types and interfaces on the given codec.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Info)(nil), nil)
	cdc.RegisterConcrete(hd.BIP44Params{}, "crypto/keys/hd/BIP44Params", nil)
	cdc.RegisterConcrete(localInfo{}, "crypto/keys/localInfo", nil)
}

// PubKeyFromBytes unmarshals public key bytes and returns a PubKey
func PubKeyFromBytes(pubKeyBytes []byte) (pubKey cryptotypes.PubKey, err error) {
	err = cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	return
}
