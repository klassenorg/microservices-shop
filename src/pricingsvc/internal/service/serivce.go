package service

import (
	"context"
	"pricingsvc/pkg/cartclient"
	"pricingsvc/pkg/catalogclient"
)

type Pricing interface {
	Calculate(ctx context.Context, id string) (int, error)
}

type Services struct {
	Pricing Pricing
}

type Deps struct {
	CartClient    *cartclient.Client
	CatalogClient *catalogclient.Client
}

func NewServices(deps Deps) *Services {
	pricingService := NewPricingService(deps.CartClient, deps.CatalogClient)
	return &Services{Pricing: pricingService}
}
