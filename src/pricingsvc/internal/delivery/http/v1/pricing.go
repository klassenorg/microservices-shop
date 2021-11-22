package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initPricingRoutes(api *gin.RouterGroup) {
	api.Use(h.hasUserIDCookie())
	api.POST("/calculate", h.calculate)
}

func (h *Handler) calculate(c *gin.Context) {
	userID, _ := c.Cookie("USER_ID")
	h.logger.Debug("getCart method started")

	price, err := h.services.Pricing.Calculate(c.Request.Context(), userID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_price": price})
}
