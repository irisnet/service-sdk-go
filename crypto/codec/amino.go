package codec

import (
	tmsr25519 "github.com/tendermint/tendermint/crypto/sr25519"

	"github.com/irisnet/service-sdk-go/codec"
	"github.com/irisnet/service-sdk-go/crypto/keys/ed25519"
	"github.com/irisnet/service-sdk-go/crypto/keys/multisig"
	"github.com/irisnet/service-sdk-go/crypto/keys/secp256k1"
	"github.com/irisnet/service-sdk-go/crypto/keys/sm2"
	cryptotypes "github.com/irisnet/service-sdk-go/crypto/types"
)

var amino *codec.LegacyAmino

func init() {
	amino = codec.NewLegacyAmino()
	RegisterCrypto(amino)
}

// RegisterCrypto registers all crypto dependency types with the provided Amino
// codec.
func RegisterCrypto(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*cryptotypes.PubKey)(nil), nil)
	cdc.RegisterConcrete(&tmsr25519.PubKey{}, tmsr25519.PubKeyName, nil)
	cdc.RegisterConcrete(&ed25519.PubKey{}, ed25519.PubKeyName, nil)
	cdc.RegisterConcrete(&secp256k1.PubKey{}, secp256k1.PubKeyName, nil)
	cdc.RegisterConcrete(&sm2.PubKey{}, sm2.PubKeyName, nil)
	cdc.RegisterConcrete(&multisig.LegacyAminoPubKey{}, multisig.PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*cryptotypes.PrivKey)(nil), nil)
	cdc.RegisterConcrete(&ed25519.PrivKey{}, ed25519.PrivKeyName, nil)
	cdc.RegisterConcrete(&tmsr25519.PrivKey{}, tmsr25519.PrivKeyName, nil)
	cdc.RegisterConcrete(&secp256k1.PrivKey{}, secp256k1.PrivKeyName, nil)
	cdc.RegisterConcrete(&sm2.PrivKey{}, sm2.PrivKeyName, nil)

}

// PrivKeyFromBytes unmarshals private key bytes and returns a PrivKey
func PrivKeyFromBytes(privKeyBytes []byte) (privKey cryptotypes.PrivKey, err error) {
	err = amino.UnmarshalBinaryBare(privKeyBytes, &privKey)
	return
}

// PubKeyFromBytes unmarshals public key bytes and returns a PubKey
func PubKeyFromBytes(pubKeyBytes []byte) (pubKey cryptotypes.PubKey, err error) {
	err = amino.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	return
}

func MarshalPubkey(pubkey cryptotypes.PubKey) []byte {
	return amino.MustMarshalBinaryBare(pubkey)
}

func MarshalPrivKey(privKey cryptotypes.PrivKey) []byte {
	return amino.MustMarshalBinaryBare(privKey)
}
