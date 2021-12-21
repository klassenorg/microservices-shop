package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "purchasesvc/gen/proto"
	"purchasesvc/internal/domain"
	"strconv"
)

func (h *Handler) Purchase(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {

	order := domain.Order{
		UserID:     req.UserId,
		FullName:   req.FullName,
		Address:    req.Address,
		CardNumber: req.CardNum,
		CVC:        req.Cvc,
		CardExp:    req.CardExp,
	}

	order, err := h.services.Purchase.CreateOrder(ctx, order)
	if err != nil {
		return &pb.PurchaseResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.PurchaseResponse{OrderId: strconv.Itoa(order.ID)}, nil
}
