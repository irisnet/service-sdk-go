package bank

import (
	"github.com/irisnet/service-sdk-go/codec"
	"github.com/irisnet/service-sdk-go/codec/types"

	sdk "github.com/irisnet/service-sdk-go/types"
)

type bankClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) BankI {
	return bankClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (b bankClient) RegisterCodec(cdc *codec.Codec) {
	registerCodec(cdc)
}

func (b bankClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)
}
