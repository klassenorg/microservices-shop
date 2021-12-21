package grpc

import (
	"context"
	pb "pricingsvc/gen/proto"
	"strconv"
)

func (h *Handler) Calculate(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	price, err := h.services.Pricing.Calculate(ctx, req.Id)
	if err != nil {
		h.logger.Errorw("error calculating cart",
			"grpc", true,
			"user_id", req.Id,
			"error", err)
		return &pb.CalculateResponse{}, err
	}

	return &pb.CalculateResponse{Price: strconv.Itoa(price)}, nil
}
