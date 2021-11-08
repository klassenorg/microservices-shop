package grpc

import (
	pb "catalogsvc/gen/proto"
	"catalogsvc/internal/service"
	"catalogsvc/pkg/logger"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
	pb.UnimplementedCatalogServiceServer
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
