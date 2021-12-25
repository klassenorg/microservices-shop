package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"purchasesvc/internal/domain"
)

type Repositories struct {
	Orders Orders
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{Orders: NewOrdersRepo(db)}
}

type Orders interface {
	Create(ctx context.Context, order domain.Order, items map[string]string) (domain.Order, error)
	Get(ctx context.Context, orderID string) (domain.Order, error)
}
