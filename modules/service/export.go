package service

import sdk "github.com/irisnet/service-sdk-go/types"

// ServiceI defines a set of interfaces in the service module
type ServiceI interface {
	sdk.Module
	//Tx
	//Query
}
