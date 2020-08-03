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
	registerCodec(cdc)
}

func (s serviceClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
}

func (s serviceClient) QueryServiceDefinition(serviceName string) (QueryServiceDefinitionResponse, sdk.Error) {
	param := struct {
		ServiceName string
	}{
		ServiceName: serviceName,
	}

	var definition ServiceDefinition
	if err := s.QueryWithResponse("custom/service/definition", param, &definition); err != nil {
		return QueryServiceDefinitionResponse{}, sdk.Wrap(err)
	}
	return definition.Convert().(QueryServiceDefinitionResponse), nil
}

func (s serviceClient) QueryServiceBinding(serviceName string, provider sdk.AccAddress) (QueryServiceBindingResponse, sdk.Error) {
	param := struct {
		ServiceName string
		Provider    sdk.AccAddress
	}{
		ServiceName: serviceName,
		Provider:    provider,
	}

	var binding ServiceBinding
	if err := s.QueryWithResponse("custom/service/binding", param, &binding); err != nil {
		return QueryServiceBindingResponse{}, sdk.Wrap(err)
	}
	return binding.Convert().(QueryServiceBindingResponse), nil
}

// QueryBindings returns all bindings of the specified service
func (s serviceClient) QueryServiceBindings(serviceName string) ([]QueryServiceBindingResponse, sdk.Error) {
	param := struct {
		ServiceName string
	}{
		ServiceName: serviceName,
	}

	var bindings serviceBindings
	if err := s.QueryWithResponse("custom/service/bindings", param, &bindings); err != nil {
		return nil, sdk.Wrap(err)
	}
	return bindings.Convert().([]QueryServiceBindingResponse), nil
}

// QueryRequest returns  the active request of the specified requestID
func (s serviceClient) QueryServiceRequest(requestID string) (QueryServiceRequestResponse, sdk.Error) {
	param := struct {
		RequestID []byte
	}{
		RequestID: sdk.MustHexBytesFrom(requestID),
	}

	var request Request
	if err := s.QueryWithResponse("custom/service/request", param, &request); request.Empty() {
		// TODO if request.Empty()
		//request, err = s.queryRequestByTxQuery(requestID)
		if err != nil {
			return QueryServiceRequestResponse{}, sdk.Wrap(err)
		}
	}
	return request.Convert().(QueryServiceRequestResponse), nil
}

// QueryRequest returns all the active requests of the specified service binding
func (s serviceClient) QueryServiceRequests(serviceName string, provider sdk.AccAddress) ([]QueryServiceRequestResponse, sdk.Error) {
	param := struct {
		ServiceName string
		Provider    sdk.AccAddress
	}{
		ServiceName: serviceName,
		Provider:    provider,
	}

	var rs requests
	if err := s.QueryWithResponse("custom/service/requests", param, &rs); err != nil {
		return nil, sdk.Wrap(err)
	}
	return rs.Convert().([]QueryServiceRequestResponse), nil
}

// QueryRequestsByReqCtx returns all requests of the specified request context ID and batch counter
func (s serviceClient) QueryRequestsByReqCtx(reqCtxID string, batchCounter uint64) ([]QueryServiceRequestResponse, sdk.Error) {
	param := struct {
		RequestContextID sdk.HexBytes
		BatchCounter     uint64
	}{
		RequestContextID: sdk.MustHexBytesFrom(reqCtxID),
		BatchCounter:     batchCounter,
	}

	var rs requests
	if err := s.QueryWithResponse("custom/service/requests_by_ctx", param, &rs); err != nil {
		return nil, sdk.Wrap(err)
	}
	return rs.Convert().([]QueryServiceRequestResponse), nil
}

// QueryResponse returns a response with the speicified request ID
func (s serviceClient) QueryServiceResponse(requestID string) (QueryServiceResponseResponse, sdk.Error) {
	param := struct {
		RequestID string
	}{
		RequestID: requestID,
	}

	var response Response
	if err := s.QueryWithResponse("custom/service/response", param, &response); response.Empty() {
		// TODO if response.Empty()
		//response, err = s.queryResponseByTxQuery(requestID)
		if err != nil {
			return QueryServiceResponseResponse{}, sdk.Wrap(nil)
		}
	}
	return response.Convert().(QueryServiceResponseResponse), nil
}

// QueryResponses returns all responses of the specified request context and batch counter
func (s serviceClient) QueryServiceResponses(reqCtxID string, batchCounter uint64) ([]QueryServiceResponseResponse, sdk.Error) {
	param := struct {
		RequestContextID sdk.HexBytes
		BatchCounter     uint64
	}{
		RequestContextID: sdk.MustHexBytesFrom(reqCtxID),
		BatchCounter:     batchCounter,
	}
	var rs responses
	if err := s.QueryWithResponse("custom/service/responses", param, &rs); err != nil {
		return nil, sdk.Wrap(err)
	}
	return rs.Convert().([]QueryServiceResponseResponse), nil
}

// QueryRequestContext return the specified request context
func (s serviceClient) QueryRequestContext(reqCtxID string) (QueryRequestContextResponse, sdk.Error) {
	param := struct {
		RequestContextID sdk.HexBytes
	}{
		RequestContextID: sdk.MustHexBytesFrom(reqCtxID),
	}

	var reqCtx RequestContext
	if err := s.QueryWithResponse("custom/service/context", param, &reqCtx); reqCtx.Empty() {
		// TODO if reqCtx.Empty()
		//reqCtx, err = s.queryRequestContextByTxQuery(reqCtxID)
		if err != nil {
			return QueryRequestContextResponse{}, sdk.Wrap(err)
		}
	}
	return reqCtx.Convert().(QueryRequestContextResponse), nil
}

//QueryFees return the earned fees for a provider
func (s serviceClient) QueryFees(provider string) (sdk.Coins, sdk.Error) {
	address, err := sdk.AccAddressFromBech32(provider)
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	param := struct {
		Provider sdk.AccAddress
	}{
		Provider: address,
	}

	bz, e := s.Query("custom/service/fees", param)
	if e != nil {
		return nil, sdk.Wrap(err)
	}

	var fee sdk.Coins
	if err := s.UnmarshalJSON(bz, &fee); err != nil {
		return nil, sdk.Wrap(err)
	}
	return fee, nil
}

func (s serviceClient) QueryParams() (QueryParamsResponse, sdk.Error) {
	var param Params
	if err := s.BaseClient.QueryParams(ModuleName, &param); err != nil {
		return QueryParamsResponse{}, err
	}
	return param.Convert().(QueryParamsResponse), nil
}
