package grpc

import (
	pb "cartsvc/gen/proto"
	"context"
)

func (h *Handler) GetCart(ctx context.Context, req *pb.CartRequest) (*pb.CartResponse, error) {
	cart, err := h.services.Cart.GetCart(ctx, req.Id)
	if err != nil {
		h.logger.Errorw("error getting cart",
			"grpc", true,
			"user_id", req.Id,
			"error", err)
		return &pb.CartResponse{}, err
	}

	return &pb.CartResponse{Cart: cart}, nil
}

func (h *Handler) AddToCart(ctx context.Context, req *pb.CartUpdateRequest) (*pb.Empty, error) {
	err := h.services.Cart.AddToCart(ctx, req.UserId, req.ProductId, int64(req.Count))
	if err != nil {
		h.logger.Errorw("error adding item to cart",
			"grpc", true,
			"user_id", req.UserId,
			"product_id", req.ProductId,
			"count", req.Count,
			"error", err)
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (h *Handler) RemoveFromCart(ctx context.Context, req *pb.CartUpdateRequest) (*pb.Empty, error) {
	err := h.services.Cart.RemoveFromCart(ctx, req.UserId, req.ProductId, int64(req.Count))
	if err != nil {
		h.logger.Errorw("error removing item from cart",
			"grpc", true,
			"user_id", req.UserId,
			"product_id", req.ProductId,
			"count", req.Count,
			"error", err)
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (h *Handler) RemoveAllFromCart(ctx context.Context, req *pb.CartRequest) (*pb.Empty, error) {
	err := h.services.Cart.RemoveAllFromCart(ctx, req.Id)
	if err != nil {
		h.logger.Errorw("error cleaning cart",
			"grpc", true,
			"user_id", req.Id,
			"error", err)
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}
