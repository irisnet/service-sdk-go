package token

import (
	sdk "github.com/irisnet/service-sdk-go/types"
)

const (
	ModuleName = "token"
)

// GetSymbol implements exported.TokenI
func (t Token) GetSymbol() string {
	return t.Symbol
}

// GetName implements exported.TokenI
func (t Token) GetName() string {
	return t.Name
}

// GetScale implements exported.TokenI
func (t Token) GetScale() uint32 {
	return t.Scale
}

// GetMinUnit implements exported.TokenI
func (t Token) GetMinUnit() string {
	return t.MinUnit
}

// GetInitialSupply implements exported.TokenI
func (t Token) GetInitialSupply() uint64 {
	return t.InitialSupply
}

// GetMaxSupply implements exported.TokenI
func (t Token) GetMaxSupply() uint64 {
	return t.MaxSupply
}

// GetMintable implements exported.TokenI
func (t Token) GetMintable() bool {
	return t.Mintable
}

// GetOwner implements exported.TokenI
func (t Token) GetOwner() sdk.AccAddress {
	return t.Owner
}

func (t Token) Convert() interface{} {
	return sdk.Token{
		Symbol:        t.Symbol,
		Name:          t.Name,
		Scale:         t.Scale,
		MinUnit:       t.MinUnit,
		InitialSupply: t.InitialSupply,
		MaxSupply:     t.MaxSupply,
		Mintable:      t.Mintable,
		Owner:         t.Owner.String(),
	}
}

type TokenInterface interface {
	GetSymbol() string
	GetName() string
	GetScale() uint32
	GetMinUnit() string
	GetInitialSupply() uint64
	GetMaxSupply() uint64
	GetMintable() bool
	GetOwner() sdk.AccAddress
}
