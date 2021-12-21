package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initCartRoutes(api *gin.RouterGroup) {
	api.Use(h.hasUserIDCookie())
	api.GET("/", h.getCart)
	api.POST("/add", h.addToCart)
	api.POST("/remove", h.removeFromCart)
	api.POST("/remove/all", h.cleanCart)
}

func (h *Handler) getCart(c *gin.Context) {
	userID, _ := c.Cookie("USER_ID")

	cart, err := h.services.Cart.GetCart(c.Request.Context(), userID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

type cartInput struct {
	ProductID string `json:"product_id" binding:"required"`
	Count     int64  `json:"count"`
}

func (h *Handler) addToCart(c *gin.Context) {
	userID, _ := c.Cookie("USER_ID")

	var inp cartInput
	if err := c.BindJSON(&inp); err != nil {
		if inp.ProductID == "" {
			h.newBadRequestResponse(c, http.StatusBadRequest, "empty product_id field")
		} else {
			h.newBadRequestResponse(c, http.StatusBadRequest, "error validating request body")
		}
		return
	}

	if err := h.services.Cart.AddToCart(c.Request.Context(), userID, inp.ProductID, inp.Count); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) removeFromCart(c *gin.Context) {
	userID, _ := c.Cookie("USER_ID")

	var inp cartInput
	if err := c.BindJSON(&inp); err != nil {
		if inp.ProductID == "" {
			h.newBadRequestResponse(c, http.StatusBadRequest, "empty product_id field")
		} else {
			h.newBadRequestResponse(c, http.StatusBadRequest, fmt.Sprintf("error validating request body %v", err))
		}
		return
	}

	if err := h.services.Cart.RemoveFromCart(c.Request.Context(), userID, inp.ProductID, inp.Count); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) cleanCart(c *gin.Context) {
	userID, _ := c.Cookie("USER_ID")

	if err := h.services.Cart.RemoveAllFromCart(c.Request.Context(), userID); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.Status(http.StatusOK)
}
