package service

import (
	"catalogsvc/internal/domain"
	"catalogsvc/internal/repository"
	"context"
)

type Products interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, id int) (domain.Product, error)
}

type Services struct {
	Products Products
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	productsService := NewProductsService(deps.Repos.Products)
	return &Services{Products: productsService}
}
