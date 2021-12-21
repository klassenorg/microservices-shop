package service

import (
	"context"
	"purchasesvc/internal/domain"
	"purchasesvc/internal/repository"
	"purchasesvc/pkg/cartclient"
	"purchasesvc/pkg/paymentclient"
	"purchasesvc/pkg/pricingclient"
)

type Purchase interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}

type Services struct {
	Purchase Purchase
}

func NewServices(deps Deps) *Services {
	purchaseService := NewPurchaseService(deps.Repos, deps.CartClient, deps.PricingClient, deps.PaymentProvider)
	return &Services{Purchase: purchaseService}
}

type Deps struct {
	Repos           *repository.Repositories
	CartClient      *cartclient.Client
	PricingClient   *pricingclient.Client
	PaymentProvider *paymentclient.Client
}
