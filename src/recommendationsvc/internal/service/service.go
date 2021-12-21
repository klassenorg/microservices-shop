package service

import (
	"context"
	"recommendationsvc/internal/domain"
	"recommendationsvc/pkg/catalogclient"
)

type Recommendation interface {
	GetRecommendations(ctx context.Context, count int) ([]domain.Product, error)
}

type Services struct {
	Recommendation Recommendation
}

type Deps struct {
	CatalogClient *catalogclient.Client
}

func NewServices(deps Deps) *Services {
	recommendationService := NewRecommendationService(deps.CatalogClient)
	return &Services{Recommendation: recommendationService}
}
