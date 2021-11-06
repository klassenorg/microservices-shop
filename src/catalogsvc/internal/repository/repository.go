package repository

import (
	"catalogsvc/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Products Products
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{Products: NewProductsRepo(db)}
}

type Products interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, id int) (domain.Product, error)
}
