package http

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "pricingsvc/internal/delivery/http/v1"
	"pricingsvc/internal/service"
	"pricingsvc/pkg/logger"
	"time"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	desugaredLogger := h.logger.SugaredLogger.Desugar()
	router.Use(
		ginzap.Ginzap(desugaredLogger, time.RFC3339, true),
		ginzap.RecoveryWithZap(desugaredLogger, true),
	)

	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.logger)

	api := router.Group("/pricing")
	{
		handlerV1.Init(api)
	}
}
