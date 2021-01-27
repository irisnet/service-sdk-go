package secp256k1

import (
	"io"
	"testing"

	"github.com/irisnet/service-sdk-go/crypto/keys/internal/benchmarking"
	"github.com/irisnet/service-sdk-go/crypto/types"
)

func BenchmarkKeyGeneration(b *testing.B) {
	benchmarkKeygenWrapper := func(reader io.Reader) types.PrivKey {
		priv := genPrivKey(reader)
		return &PrivKey{Key: priv}
	}
	benchmarking.BenchmarkKeyGeneration(b, benchmarkKeygenWrapper)
}

func BenchmarkSigning(b *testing.B) {
	priv := GenPrivKey()
	benchmarking.BenchmarkSigning(b, priv)
}

func BenchmarkVerification(b *testing.B) {
	priv := GenPrivKey()
	benchmarking.BenchmarkVerification(b, priv)
}
