package service

import "github.com/irisnet/service-sdk-go/codec"

const (
	// ModuleName define module name
	ModuleName = "service"
)

var (
	RequestContextBatchStateToStringMap = map[RequestContextBatchState]string{
		BATCHRUNNING:   "running",
		BATCHCOMPLETED: "completed",
	}

	RequestContextStateToStringMap = map[RequestContextState]string{
		RUNNING:   "running",
		PAUSED:    "paused",
		COMPLETED: "completed",
	}
)

//==========================================for QueryWithResponse==========================================

func (r ServiceDefinition) Convert() interface{} {
	return QueryServiceDefinitionResponse{
		Name:              r.Name,
		Description:       r.Description,
		Tags:              r.Tags,
		Author:            r.Author,
		AuthorDescription: r.AuthorDescription,
		Schemas:           r.Schemas,
	}
}

func (b ServiceBinding) Convert() interface{} {
	return QueryServiceBindingResponse{
		ServiceName:  b.ServiceName,
		Provider:     b.Provider,
		Deposit:      b.Deposit,
		Pricing:      b.Pricing,
		QoS:          b.QoS,
		Available:    b.Available,
		DisabledTime: b.DisabledTime,
		Owner:        b.Owner,
	}
}

type serviceBindings []ServiceBinding

func (bs serviceBindings) Convert() interface{} {
	bindings := make([]QueryServiceBindingResponse, len(bs))
	for i, binding := range bs {
		bindings[i] = binding.Convert().(QueryServiceBindingResponse)
	}
	return bindings
}

func (r Request) Empty() bool {
	return len(r.ServiceName) == 0
}

func (r Request) Convert() interface{} {
	return QueryServiceRequestResponse{
		ID:                         r.ID.String(),
		ServiceName:                r.ServiceName,
		Provider:                   r.Provider,
		Consumer:                   r.Consumer,
		Input:                      r.Input,
		ServiceFee:                 r.ServiceFee,
		SuperMode:                  r.SuperMode,
		RequestHeight:              r.RequestHeight,
		ExpirationHeight:           r.ExpirationHeight,
		RequestContextID:           r.RequestContextID.String(),
		RequestContextBatchCounter: r.RequestContextBatchCounter,
	}
}

type requests []Request

func (rs requests) Convert() interface{} {
	requests := make([]QueryServiceRequestResponse, len(rs))
	for i, request := range rs {
		requests[i] = request.Convert().(QueryServiceRequestResponse)
	}
	return requests
}

func (r Response) Empty() bool {
	return len(r.Provider) == 0
}

func (r Response) Convert() interface{} {
	return QueryServiceResponseResponse{
		Provider:                   r.Provider,
		Consumer:                   r.Consumer,
		Output:                     r.Output,
		Result:                     r.Result,
		RequestContextID:           r.RequestContextID.String(),
		RequestContextBatchCounter: r.RequestContextBatchCounter,
	}
}

type responses []Response

func (rs responses) Convert() interface{} {
	responses := make([]QueryServiceResponseResponse, len(rs))
	for i, response := range rs {
		responses[i] = response.Convert().(QueryServiceResponseResponse)
	}
	return responses
}

func (state RequestContextBatchState) String() string {
	return RequestContextBatchStateToStringMap[state]
}

func (state RequestContextState) String() string {
	return RequestContextStateToStringMap[state]
}

// Empty returns true if empty
func (r RequestContext) Empty() bool {
	return len(r.ServiceName) == 0
}

func (r RequestContext) Convert() interface{} {
	return QueryRequestContextResponse{
		ServiceName:        r.ServiceName,
		Providers:          r.Providers,
		Consumer:           r.Consumer,
		Input:              r.Input,
		ServiceFeeCap:      r.ServiceFeeCap,
		Timeout:            r.Timeout,
		SuperMode:          r.SuperMode,
		Repeated:           r.Repeated,
		RepeatedFrequency:  r.RepeatedFrequency,
		RepeatedTotal:      r.RepeatedTotal,
		BatchCounter:       r.BatchCounter,
		BatchRequestCount:  r.BatchRequestCount,
		BatchResponseCount: r.BatchResponseCount,
		BatchState:         r.BatchState.String(),
		State:              r.State.String(),
		ResponseThreshold:  r.ResponseThreshold,
		ModuleName:         r.ModuleName,
	}
}
func (p Params) Convert() interface{} {
	return QueryParamsResponse{
		MaxRequestTimeout:    p.MaxRequestTimeout,
		MinDepositMultiple:   p.MinDepositMultiple,
		MinDeposit:           p.MinDeposit.String(),
		ServiceFeeTax:        p.ServiceFeeTax.String(),
		SlashFraction:        p.SlashFraction.String(),
		ComplaintRetrospect:  p.ComplaintRetrospect,
		ArbitrationTimeLimit: p.ArbitrationTimeLimit,
		TxSizeLimit:          p.TxSizeLimit,
		BaseDenom:            p.BaseDenom,
	}
}

func registerCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgDefineService{}, "irismod/service/MsgDefineService", nil)
	cdc.RegisterConcrete(MsgBindService{}, "irismod/service/MsgBindService", nil)
	cdc.RegisterConcrete(MsgUpdateServiceBinding{}, "irismod/service/MsgUpdateServiceBinding", nil)
	cdc.RegisterConcrete(MsgSetWithdrawAddress{}, "irismod/service/MsgSetWithdrawAddress", nil)
	cdc.RegisterConcrete(MsgDisableServiceBinding{}, "irismod/service/MsgDisableServiceBinding", nil)
	cdc.RegisterConcrete(MsgEnableServiceBinding{}, "irismod/service/MsgEnableServiceBinding", nil)
	cdc.RegisterConcrete(MsgRefundServiceDeposit{}, "irismod/service/MsgRefundServiceDeposit", nil)
	cdc.RegisterConcrete(MsgCallService{}, "irismod/service/MsgCallService", nil)
	cdc.RegisterConcrete(MsgRespondService{}, "irismod/service/MsgRespondService", nil)
	cdc.RegisterConcrete(MsgPauseRequestContext{}, "irismod/service/MsgPauseRequestContext", nil)
	cdc.RegisterConcrete(MsgStartRequestContext{}, "irismod/service/MsgStartRequestContext", nil)
	cdc.RegisterConcrete(MsgKillRequestContext{}, "irismod/service/MsgKillRequestContext", nil)
	cdc.RegisterConcrete(MsgUpdateRequestContext{}, "irismod/service/MsgUpdateRequestContext", nil)
	cdc.RegisterConcrete(MsgWithdrawEarnedFees{}, "irismod/service/MsgWithdrawEarnedFees", nil)
}
