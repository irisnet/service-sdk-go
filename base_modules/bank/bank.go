package bank

import (
	"context"

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

func (b bankClient) Name() string {
	return ModuleName
}

func (b bankClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

// QueryAccount returns account information by the specified address
func (b bankClient) QueryAccount(address string) (sdk.BaseAccount, sdk.Error) {
	conn, err := b.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return sdk.BaseAccount{}, sdk.Wrap(err)
	}

	account, err := b.BaseClient.QueryAccount(address)
	if err != nil {
		return sdk.BaseAccount{}, sdk.Wrap(err)
	}

	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.BaseAccount{}, sdk.Wrap(err)
	}

	request := &QueryAllBalancesRequest{
		Address:    addr,
		Pagination: nil,
	}
	balances, err := NewQueryClient(conn).AllBalances(context.Background(), request)
	if err != nil {
		return sdk.BaseAccount{}, sdk.Wrap(err)
	}

	account.Coins = balances.Balances
	return account, nil
}
