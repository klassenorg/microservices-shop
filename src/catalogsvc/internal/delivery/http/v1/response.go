package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	h.logger.Error(message)
	c.AbortWithStatusJSON(statusCode, gin.H{"error":message})
}