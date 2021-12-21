package grpc

import (
	pb "purchasesvc/gen/proto"
	"purchasesvc/internal/service"
	"purchasesvc/pkg/logger"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
	pb.UnimplementedPurchaseServiceServer
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
