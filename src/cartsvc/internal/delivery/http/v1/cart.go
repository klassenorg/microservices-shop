package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initCartRoutes(api *gin.RouterGroup) {
	api.GET("/", h.getCart)
	api.POST("/add", h.addToCart)
	api.POST("/remove", h.removeFromCart)
	api.POST("/remove/all", h.cleanCart)
}

func (h *Handler) getCart(c *gin.Context) {
	userID, err := c.Cookie("USER_ID")
	if err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cart, err := h.services.Cart.GetCart(c.Request.Context(), userID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

type cartInput struct {
	ProductID string `json:"product_id" binding:"required"`
	Count     int64  `json:"count"`
}

func (h *Handler) addToCart(c *gin.Context) {
	userID, err := c.Cookie("USER_ID")
	if err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp cartInput
	inp.Count = 1
	if err := c.BindJSON(&inp); err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Cart.AddToCart(c.Request.Context(), userID, inp.ProductID, inp.Count); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) removeFromCart(c *gin.Context) {
	userID, err := c.Cookie("USER_ID")
	if err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp cartInput
	inp.Count = 1
	if err := c.BindJSON(&inp); err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Cart.RemoveFromCart(c.Request.Context(), userID, inp.ProductID, inp.Count); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) cleanCart(c *gin.Context) {
	userID, err := c.Cookie("USER_ID")
	if err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Cart.RemoveAllFromCart(c.Request.Context(), userID); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
