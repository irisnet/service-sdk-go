// Package token allows individuals and companies to create and issue their own tokens.
//

package token

import (
	"github.com/irisnet/service-sdk-go/codec"
	"github.com/irisnet/service-sdk-go/codec/types"
	sdk "github.com/irisnet/service-sdk-go/types"
)

type tokenClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) TokenI {
	return tokenClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (t tokenClient) Name() string {
	return ModuleName
}

func (t tokenClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (t tokenClient) QueryToken(denom string) (sdk.Token, error) {
	return t.BaseClient.QueryToken(denom)
}
