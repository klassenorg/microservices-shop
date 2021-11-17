package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) hasUserIDCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := c.Cookie("USER_ID")
		if err != nil {
			h.newBadRequestResponse(c, http.StatusBadRequest, "USER_ID cookie not present")
			return
		}
		if userID == "" {
			h.newBadRequestResponse(c, http.StatusBadRequest, "USER_ID cookie is empty")
			return
		}
		c.Next()
	}
}
