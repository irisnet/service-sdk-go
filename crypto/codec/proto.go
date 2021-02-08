package codec

import (
	codectypes "github.com/irisnet/service-sdk-go/codec/types"
	"github.com/irisnet/service-sdk-go/crypto/keys/ed25519"
	"github.com/irisnet/service-sdk-go/crypto/keys/multisig"
	"github.com/irisnet/service-sdk-go/crypto/keys/secp256k1"
	"github.com/irisnet/service-sdk-go/crypto/keys/sm2"
	cryptotypes "github.com/irisnet/service-sdk-go/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.crypto.Pubkey", (*cryptotypes.PubKey)(nil))
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &sm2.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})
}
