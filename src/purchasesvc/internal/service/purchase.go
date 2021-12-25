package service

import (
	"context"
	"errors"
	"purchasesvc/internal/domain"
	"purchasesvc/internal/repository"
	"purchasesvc/pkg/cartclient"
	"purchasesvc/pkg/paymentclient"
	"purchasesvc/pkg/pricingclient"
)

type PurchaseService struct {
	Repos           *repository.Repositories
	CartClient      *cartclient.Client
	PricingClient   *pricingclient.Client
	PaymentProvider *paymentclient.Client
}

func NewPurchaseService(repos *repository.Repositories, cartClient *cartclient.Client, pricingClient *pricingclient.Client, paymentProvider *paymentclient.Client) *PurchaseService {
	return &PurchaseService{Repos: repos, CartClient: cartClient, PricingClient: pricingClient, PaymentProvider: paymentProvider}
}

func (s *PurchaseService) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {

	items, err := s.CartClient.GetCart(ctx, order.UserID)
	if err != nil {
		return domain.Order{}, err
	}

	if len(items) < 1 {
		return domain.Order{}, errors.New("empty cart")
	}

	order.TotalPrice, err = s.PricingClient.Calculate(ctx, order.UserID)
	if err != nil {
		return domain.Order{}, err
	}

	if err := s.PaymentProvider.Pay(order.CardNumber, order.CVC, order.CardExp); err != nil {
		return domain.Order{}, err
	}

	order, err = s.Repos.Orders.Create(ctx, order, items)
	if err != nil {
		return domain.Order{}, err
	}

	//TODO: handle somehow(maybe inject logger to service and log error or made custom error)
	_ = s.CartClient.RemoveAllFromCart(ctx, order.UserID)

	return order, nil
}

func (s *PurchaseService) GetOrder(ctx context.Context, orderID string) (domain.Order, error) {
	return s.Repos.Orders.Get(ctx, orderID)

}
