package grpc

import (
	"context"
	pb "recommendationsvc/gen/proto"
)

func (h *Handler) GetRecommendations(ctx context.Context, req *pb.GetRecommendationsRequest) (*pb.Products, error) {
	recommendations, err := h.services.Recommendation.GetRecommendations(ctx, int(req.Count))
	if err != nil {
		return &pb.Products{}, err
	}

	out := new(pb.Products)

	for _, product := range recommendations {
		out.Products = append(out.Products, &pb.Product{
			Id:          uint32(product.ProductID),
			Name:        product.Name,
			Description: product.Description,
			Price:       uint32(product.Price),
			ImagePath:   product.ImagePath,
		})
	}

	return out, nil
}
