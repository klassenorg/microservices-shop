package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initRecommendationsRoutes(api *gin.RouterGroup) {
	products := api.Group("/")
	{
		products.GET("/:count", h.getRecommendations)
	}
}

func (h *Handler) getRecommendations(c *gin.Context) {
	countParam := c.Param("count")
	count, err := strconv.Atoi(countParam)
	if err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, "wrong count param")
		return
	}

	res, err := h.services.Recommendation.GetRecommendations(c.Request.Context(), count)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.JSON(http.StatusOK, res)
}
