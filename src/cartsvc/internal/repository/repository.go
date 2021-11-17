package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repositories struct {
	Cart Cart
}

func NewRepositories(db *redis.Client) *Repositories {
	return &Repositories{Cart: NewCartRepo(db)}
}

type Cart interface {
	GetByID(ctx context.Context, id string) (map[string]string, error)
	AddByID(ctx context.Context, id string, product string, count int64) error
	RemoveByID(ctx context.Context, id string, product string, count int64) error
	RemoveAllByID(ctx context.Context, id string) error
}
