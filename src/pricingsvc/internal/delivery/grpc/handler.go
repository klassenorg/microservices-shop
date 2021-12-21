package grpc

import (
	pb "pricingsvc/gen/proto"
	"pricingsvc/internal/service"
	"pricingsvc/pkg/logger"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
	pb.UnimplementedPricingServiceServer
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
