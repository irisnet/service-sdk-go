package sm2

import (
	"bytes"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"io"
	"math/big"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tjfoc/gmsm/sm2"
)

const (
	PrivKeyAminoName = "tendermint/PrivKeySm2"
	PubKeyAminoName  = "tendermint/PubKeySm2"

	PrivKeySize   = 32
	PubKeySize    = 33
	SignatureSize = 64
)

type PrivKeySm2 [PrivKeySize]byte
type PubKeySm2 [PubKeySize]byte

var _ crypto.PrivKey = PrivKeySm2{}
var _ crypto.PubKey = PubKeySm2{}

// --------------------------------------------------------

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(PubKeySm2{}, PubKeyAminoName, nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(PrivKeySm2{}, PrivKeyAminoName, nil)
}

// --------------------------------------------------------

func (privKey PrivKeySm2) Bytes() []byte {
	return cdc.MustMarshalBinaryBare(privKey)
}

func (privKey PrivKeySm2) Sign(msg []byte) ([]byte, error) {
	priv := privKey.GetPrivateKey()
	r, s, err := sm2.Sm2Sign(priv, msg, nil)
	if err != nil {
		return nil, err
	}
	R := r.Bytes()
	S := s.Bytes()
	sig := make([]byte, 64)
	copy(sig[32-len(R):32], R[:])
	copy(sig[64-len(S):64], S[:])

	return sig, nil
}

func (privKey PrivKeySm2) Sm2Sign(msg []byte) ([]byte, error) {
	r, s, err := sm2.Sm2Sign(privKey.GetPrivateKey(), msg, nil)
	if err != nil {
		return nil, err
	}

	R := r.Bytes()
	S := s.Bytes()
	sig := make([]byte, 64)
	copy(sig[32-len(R):32], R[:])
	copy(sig[64-len(S):64], S[:])

	return sig, nil
}

func (privKey PrivKeySm2) PubKey() crypto.PubKey {
	priv := privKey.GetPrivateKey()
	compPubkey := sm2.Compress(&priv.PublicKey)
	var pubKey PubKeySm2
	copy(pubKey[:], compPubkey)

	return pubKey
}

func (privKey PrivKeySm2) PubKeySm2() PubKeySm2 {
	priv := privKey.GetPrivateKey()
	compPubkey := sm2.Compress(&priv.PublicKey)
	var pubKey PubKeySm2
	copy(pubKey[:], compPubkey)

	return pubKey
}

func (privKey PrivKeySm2) Equals(other crypto.PrivKey) bool {
	if otherSm2, ok := other.(PrivKeySm2); ok {
		return subtle.ConstantTimeCompare(privKey[:], otherSm2[:]) == 1
	}

	return false
}

func (privKey PrivKeySm2) GetPrivateKey() *sm2.PrivateKey {
	k := new(big.Int).SetBytes(privKey[:32])
	c := sm2.P256Sm2()
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())

	return priv
}

func GenPrivKey() PrivKeySm2 {
	return genPrivKey(crypto.CReader())
}

func genPrivKey(rand io.Reader) PrivKeySm2 {
	seed := make([]byte, 32)
	if _, err := io.ReadFull(rand, seed); err != nil {
		panic(err)
	}

	privKey, err := sm2.GenerateKey()
	if err != nil {
		panic(err)
	}

	var privKeySm2 PrivKeySm2
	copy(privKeySm2[:], privKey.D.Bytes())

	return privKeySm2
}

func GenPrivKeySm2FromSecret(secret []byte) PrivKeySm2 {
	one := new(big.Int).SetInt64(1)
	secHash := sha256.Sum256(secret)

	k := new(big.Int).SetBytes(secHash[:])
	n := new(big.Int).Sub(sm2.P256Sm2().Params().N, one)
	k.Mod(k, n)
	k.Add(k, one)

	var privKeySm2 PrivKeySm2
	copy(privKeySm2[:], k.Bytes())

	return privKeySm2
}

// --------------------------------------------------------

func (pubKey PubKeySm2) Address() crypto.Address {
	return crypto.Address(tmhash.SumTruncated(pubKey[:]))
}

func (pubKey PubKeySm2) Bytes() []byte {
	return cdc.MustMarshalBinaryBare(pubKey)
}

func (pubKey PubKeySm2) VerifyBytes(msg []byte, sig []byte) bool {
	if len(sig) != SignatureSize {
		return false
	}

	publicKey := sm2.Decompress(pubKey[:])
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	return sm2.Sm2Verify(publicKey, msg, nil, r, s)
}

func (pubKey PubKeySm2) Sm2VerifyBytes(msg []byte, sig []byte) bool {
	if len(sig) != SignatureSize {
		return false
	}

	publicKey := sm2.Decompress(pubKey[:])
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	return sm2.Sm2Verify(publicKey, msg, nil, r, s)
}

func (pubKey PubKeySm2) String() string {
	return fmt.Sprintf("PubKeySm2{%X}", pubKey[:])
}

func (pubKey PubKeySm2) Equals(other crypto.PubKey) bool {
	if otherSm2, ok := other.(PubKeySm2); ok {
		return bytes.Equal(pubKey[:], otherSm2[:])
	} else {
		return false
	}
}
