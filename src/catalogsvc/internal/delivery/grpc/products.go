package grpc

import (
	pb "catalogsvc/gen/proto"
	"context"
)

func (h *Handler) GetAllProducts(ctx context.Context, _ *pb.Empty) (*pb.Products, error) {
	products, err := h.services.Products.GetAll(ctx)
	if err != nil {
		h.logger.Errorw("error getting products",
			"grpc", true,
			"error", err)
		return nil, err
	}

	out := make([]*pb.Product, 0, len(products))

	for _, product := range products {
		out = append(out, &pb.Product{
			Id:          uint32(product.ProductID),
			Name:        product.Name,
			Description: product.Description,
			Price:       uint32(product.Price),
			ImagePath:   product.ImagePath,
		})
	}

	return &pb.Products{Products: out}, nil
}

func (h *Handler) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.Product, error) {
	product, err := h.services.Products.GetByID(ctx, int(req.Id))
	if err != nil {
		h.logger.Errorw("error getting product by id",
			"grpc", true,
			"id", req.Id,
			"error", err)
		return nil, err
	}

	return &pb.Product{
		Id:          uint32(product.ProductID),
		Name:        product.Name,
		Description: product.Description,
		Price:       uint32(product.Price),
		ImagePath:   product.ImagePath,
	}, nil
}
