package integration_test

import (
	"github.com/irisnet/service-sdk-go/modules/service"
	sdk "github.com/irisnet/service-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) TestQuery() {
	//serviceName := "assettransfer"
	//definition, err := s.Service.QueryServiceDefinition("assettransfer")
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(),definition)

	//binding, err := s.Service.QueryServiceBinding("assettransfer",s.rootAccount.Address)
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(),binding)

	//bindings, err := s.Service.QueryServiceBindings("assettransfer")
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(),bindings)

	//requestResponse, err := s.Service.QueryServiceRequest("B17A1F4B700F7199F6CE8DDDC5CEFF2CF4D4A4792C30455FCA59E092846654BB00000000000000000000000000000001000000000000CE480000")
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), requestResponse)

	//requestResponse, err := s.Service.QueryServiceRequests(serviceName,s.rootAccount.Address)
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), requestResponse)

	//requestResponse, err := s.Service.QueryRequestsByReqCtx("ADDC93462A9839A6016D4A881D900289B8674D71BE10A06CAF737D8897B00EC50000000000000000",1)
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), requestResponse)

	//response, err := s.Service.QueryServiceResponse("A8DD02EA506C2148D9B5BDC07EE193268E1D11E81016418D01AD2C4C967DD98A00000000000000000000000000000001000000000000CF5D0000")
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), response)

	//responses, err := s.Service.QueryServiceResponses("A8DD02EA506C2148D9B5BDC07EE193268E1D11E81016418D01AD2C4C967DD98A0000000000000000",1)
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), responses)

	//requestContext, err := s.Service.QueryRequestContext("86E7B1AC2E9370053733EB67448711CF285F386D91158B225E3F781E001AC25E0000000000000000")
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), requestContext)

	//fees, err := s.Service.QueryFees(s.rootAccount.Address.String())
	//require.NoError(s.T(), err)
	//require.NotEmpty(s.T(), fees)

	params, err := s.Service.QueryParams()
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), params)
}

func (s IntegrationTestSuite) TestTx() {
	schemas := `{"input":{"type":"object"},"output":{"type":"object"},"error":{"type":"object"}}`

	baseTx := sdk.BaseTx{
		From: s.rootAccount.Name,
		Gas:  200000,
		//Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.rootAccount.Password,
	}

	definition := service.DefineServiceRequest{
		ServiceName:       "assettransfer",
		Description:       "asset transfer",
		Tags:              nil,
		AuthorDescription: "tester",
		Schemas:           schemas,
	}
	result, err := s.Service.DefineService(definition, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)
}

func (s IntegrationTestSuite) TestBind() {
	pricing := `{"price":"1stake"}`

	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      200000,
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}

	deposit, e := sdk.ParseDecCoins("20000stake")
	require.NoError(s.T(), e)

	binding := service.BindServiceRequest{
		ServiceName: "test0804",
		Deposit:     deposit,
		Pricing:     pricing,
		QoS:         1,
	}
	result, err := s.Service.BindService(binding, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)

}
