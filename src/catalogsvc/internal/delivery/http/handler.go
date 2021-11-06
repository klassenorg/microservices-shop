package http

import (
	v1 "catalogsvc/internal/delivery/http/v1"
	"catalogsvc/internal/service"
	"catalogsvc/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Services
	logger *logger.Logger
}

func NewHandler(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger: logger,
	}
}


func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		)

	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.logger)
	api := router.Group("/catalog")
	{
		handlerV1.Init(api)
	}
}