package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	h.logger.Errorw(err.Error(),
		"location", c.Request.RequestURI)
	c.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}

func (h *Handler) newBadRequestResponse(c *gin.Context, statusCode int, message string) {
	h.logger.Debugw(message,
		"location", c.Request.RequestURI)
	c.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}
