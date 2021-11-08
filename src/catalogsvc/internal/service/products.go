package service

import (
	"catalogsvc/internal/domain"
	"catalogsvc/internal/repository"
	"context"
)

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{repo: repo}
}

func (p ProductsService) GetAll(ctx context.Context) ([]domain.Product, error) {
	return p.repo.GetAll(ctx)
}

func (p ProductsService) GetByID(ctx context.Context, id int) (domain.Product, error) {
	return p.repo.GetByID(ctx, id)
}
