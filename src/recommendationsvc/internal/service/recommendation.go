package service

import (
	"context"
	"math/rand"
	"recommendationsvc/internal/domain"
	"recommendationsvc/pkg/catalogclient"
	"time"
)

type RecommendationService struct {
	catalogClient *catalogclient.Client
}

func (r *RecommendationService) GetRecommendations(ctx context.Context, count int) ([]domain.Product, error) {

	//TODO:cache products
	products, err := r.catalogClient.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	if count > len(products) {
		count = len(products)
	}

	out := make([]domain.Product, 0, count)
	rand.Seed(time.Now().Unix())
	alreadyIn := make(map[int]bool)
	var randNum = rand.Intn(len(products))
	for i := 0; i < count; i++ {
		for alreadyIn[randNum] {
			randNum = rand.Intn(len(products))
		}
		alreadyIn[randNum] = true
		out = append(out, domain.Product(products[randNum]))
	}

	return out, nil
}

func NewRecommendationService(catalogClient *catalogclient.Client) *RecommendationService {
	return &RecommendationService{catalogClient: catalogClient}
}
