package v1

import (
	"github.com/gin-gonic/gin"
	"purchasesvc/internal/service"
	"purchasesvc/pkg/logger"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initPurchaseRoutes(v1)
	}
}
