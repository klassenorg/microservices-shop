package service

import (
	"cartsvc/internal/repository"
	"context"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (c *CartService) GetCart(ctx context.Context, id string) (map[string]string, error) {
	return c.repo.GetByID(ctx, id)
}

func (c *CartService) AddToCart(ctx context.Context, id string, product string, count int64) error {
	return c.repo.AddByID(ctx, id, product, count)
}

func (c *CartService) RemoveFromCart(ctx context.Context, id string, product string, count int64) error {
	return c.repo.RemoveByID(ctx, id, product, count)
}

func (c *CartService) RemoveAllFromCart(ctx context.Context, id string) error {
	return c.repo.RemoveAllByID(ctx, id)
}
