package sdk

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/irisnet/service-sdk-go/base_modules"
	"github.com/irisnet/service-sdk-go/base_modules/bank"
	"github.com/irisnet/service-sdk-go/base_modules/token"
	"github.com/irisnet/service-sdk-go/codec"
	cdctypes "github.com/irisnet/service-sdk-go/codec/types"
	cryptocodec "github.com/irisnet/service-sdk-go/crypto/codec"
	"github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	txtypes "github.com/irisnet/service-sdk-go/types/tx"
)

// ServiceClient exports service.ServiceClient
type ServiceClient = service.ServiceClient

// NewServiceClient contructs a service client
func NewServiceClient(cfg types.ClientConfig) ServiceClient {
	return NewClient(cfg).Service.(ServiceClient)
}

type Client struct {
	logger         log.Logger
	moduleManager  map[string]types.Module
	encodingConfig types.EncodingConfig

	types.BaseClient
	Bank    bank.BankI
	Token   token.TokenI
	Service service.ServiceI
}

func NewClient(cfg types.ClientConfig) Client {
	encodingConfig := makeEncodingConfig()

	// create an instance of baseClient
	baseClient := base_modules.NewBaseClient(cfg, encodingConfig, nil)

	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
	tokenClient := token.NewClient(baseClient, encodingConfig.Marshaler)
	serviceClient := service.NewClient(baseClient, encodingConfig.Marshaler)

	client := &Client{
		logger:         baseClient.Logger(),
		BaseClient:     baseClient,
		Bank:           bankClient,
		Token:          tokenClient,
		Service:        serviceClient,
		moduleManager:  make(map[string]types.Module),
		encodingConfig: encodingConfig,
	}

	client.RegisterModule(
		bankClient,
		tokenClient,
		serviceClient,
	)

	return *client
}

func (client *Client) SetLogger(logger log.Logger) {
	client.BaseClient.SetLogger(logger)
}

func (client *Client) Codec() *codec.LegacyAmino {
	return client.encodingConfig.Amino
}

func (client *Client) AppCodec() codec.Marshaler {
	return client.encodingConfig.Marshaler
}

func (client *Client) Manager() types.BaseClient {
	return client.BaseClient
}

func (client *Client) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		_, ok := client.moduleManager[m.Name()]
		if ok {
			panic(fmt.Sprintf("%s has register", m.Name()))
		}

		// m.RegisterCodec(client.encodingConfig.Amino)
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
		client.moduleManager[m.Name()] = m
	}
}

func makeEncodingConfig() types.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, types.DefaultPublicKeyCodec{}, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}

	RegisterLegacyAminoCodec(encodingConfig.Amino)
	RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*types.Msg)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers the sdk message type.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*types.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
}
