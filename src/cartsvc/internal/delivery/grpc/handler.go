package grpc

import (
	pb "cartsvc/gen/proto"
	"cartsvc/internal/service"
	"cartsvc/pkg/logger"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
	pb.UnimplementedCartServiceServer
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
