package sdk

import (
	"io"

	"github.com/irisnet/service-sdk-go/codec"
	cdctypes "github.com/irisnet/service-sdk-go/codec/types"
	"github.com/irisnet/service-sdk-go/modules"
	"github.com/irisnet/service-sdk-go/modules/keys"
	"github.com/irisnet/service-sdk-go/modules/service"
	"github.com/irisnet/service-sdk-go/std"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/utils/log"
)

type IServiceClient struct {
	logger *log.Logger

	types.WSClient
	types.TxManager
	types.TokenConvert

	Service service.ServiceI
	Key     keys.KeyI
}

func NewIServiceClient(cfg types.ClientConfig) IServiceClient {
	//create cdc for encoding and decoding
	cdc := types.NewCodec()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	appCodec := std.NewAppCodec(cdc, interfaceRegistry)

	//create a instance of baseClient
	baseClient := modules.NewBaseClient(cfg, appCodec)

	serviceClient := service.NewClient(baseClient, appCodec)
	keysClient := keys.NewClient(baseClient, appCodec)

	client := &IServiceClient{
		logger:       baseClient.Logger(),
		WSClient:     baseClient,
		TxManager:    baseClient,
		TokenConvert: baseClient,

		//Bank:    bankClient,
		Key:     keysClient,
		Service: serviceClient,
	}

	client.RegisterModule(
		cdc,
		interfaceRegistry,
		keysClient,
		serviceClient,
	)
	return *client
}

func (s *IServiceClient) SetOutput(w io.Writer) {
	s.logger.SetOutput(w)
}

func (s *IServiceClient) RegisterModule(cdc *codec.Codec, interfaceRegistry cdctypes.InterfaceRegistry, ms ...types.Module) {
	for _, m := range ms {
		m.RegisterCodec(cdc)
		m.RegisterInterfaceTypes(interfaceRegistry)
	}
	cdc.Seal()
}
