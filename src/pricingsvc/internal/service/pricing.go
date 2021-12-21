package service

import (
	"context"
	"pricingsvc/pkg/cartclient"
	"pricingsvc/pkg/catalogclient"
	"strconv"
)

type PricingService struct {
	cartClient    *cartclient.Client
	catalogClient *catalogclient.Client
}

func NewPricingService(cartClient *cartclient.Client, catalogClient *catalogclient.Client) *PricingService {
	return &PricingService{cartClient: cartClient, catalogClient: catalogClient}
}

func (p *PricingService) Calculate(ctx context.Context, id string) (int, error) {
	cart, err := p.cartClient.GetCart(ctx, id)
	if err != nil {
		return -1, err
	}

	var total int
	for item, count := range cart {
		productID, err := strconv.Atoi(item)
		if err != nil {
			return -1, err
		}
		product, err := p.catalogClient.GetProductByID(ctx, productID)
		price := product.Price
		if err != nil {
			return -1, err
		}
		countInt, err := strconv.Atoi(count)
		if err != nil {
			return -1, err
		}
		total += price * countInt
	}

	return total, nil
}
