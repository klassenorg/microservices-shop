package grpc

import (
	pb "recommendationsvc/gen/proto"
	"recommendationsvc/internal/service"
	"recommendationsvc/pkg/logger"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
	pb.UnimplementedRecommendationServiceServer
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
