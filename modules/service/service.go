package service

import (
	"github.com/irisnet/service-sdk-go/codec"
	"github.com/irisnet/service-sdk-go/codec/types"
	sdk "github.com/irisnet/service-sdk-go/types"
)

type serviceClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) ServiceI {
	return serviceClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (s serviceClient) RegisterCodec(cdc *codec.Codec) {
}

func (s serviceClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
}
