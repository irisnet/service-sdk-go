package service

const (
	// ModuleName define module name
	ModuleName = "service"
)

func (p Params) Convert() interface{} {
	//return QueryParamsResponse{
	//	MaxRequestTimeout:    p.MaxRequestTimeout,
	//	MinDepositMultiple:   p.MinDepositMultiple,
	//	MinDeposit:           p.MinDeposit.String(),
	//	ServiceFeeTax:        p.ServiceFeeTax.String(),
	//	SlashFraction:        p.SlashFraction.String(),
	//	ComplaintRetrospect:  p.ComplaintRetrospect,
	//	ArbitrationTimeLimit: p.ArbitrationTimeLimit,
	//	TxSizeLimit:          p.TxSizeLimit,
	//	BaseDenom:            p.BaseDenom,
	//}
	return p
}
