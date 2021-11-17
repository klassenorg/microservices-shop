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

func (c *CartService) GetCart(ctx context.Context, userID string) (map[string]string, error) {
	return c.repo.GetByID(ctx, userID)
}

func (c *CartService) AddToCart(ctx context.Context, userID string, productID string, count int64) error {
	if count <= 0 {
		count = 1
	}
	return c.repo.AddByID(ctx, userID, productID, count)
}

func (c *CartService) RemoveFromCart(ctx context.Context, userID string, productID string, count int64) error {
	if count <= 0 {
		count = 1
	}
	return c.repo.RemoveByID(ctx, userID, productID, count)
}

func (c *CartService) RemoveAllFromCart(ctx context.Context, userID string) error {
	return c.repo.RemoveAllByID(ctx, userID)
}
