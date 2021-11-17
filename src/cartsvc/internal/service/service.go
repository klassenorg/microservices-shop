package service

import (
	"cartsvc/internal/repository"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Cart interface {
	GetCart(ctx context.Context, id string) (map[string]string, error)
	AddToCart(ctx context.Context, id string, product string, count int64) error
	RemoveFromCart(ctx context.Context, id string, product string, count int64) error
	RemoveAllFromCart(ctx context.Context, id string) error
}

type Services struct {
	Cart Cart
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	cartService := NewCartService(deps.Repos.Cart)
	return &Services{Cart: cartService}
}
