package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"purchasesvc/internal/domain"
)

func (h *Handler) initPurchaseRoutes(api *gin.RouterGroup) {

	purchase := api.Group("/")
	{
		purchase.Use(h.hasUserIDCookie())
		purchase.POST("/", h.purchase)
	}
}

type purchaseInput struct {
	FullName   string `json:"full_name" binding:"required,min=2,max=64"`
	Address    string `json:"address" binding:"required,min=6"`
	CardNumber string `json:"card_number" binding:"required,min=16,max=16"`
	CVC        string `json:"cvc" binding:"required,min=3,max=3"`
	CardExp    string `json:"exp" binding:"required,min=4,max=4"`
}

func (h *Handler) purchase(c *gin.Context) {
	userID, _ := c.Cookie("USER_ID")

	var input purchaseInput
	if err := c.BindJSON(&input); err != nil {
		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	order := domain.Order{
		UserID:     userID,
		FullName:   input.FullName,
		Address:    input.Address,
		CardNumber: input.CardNumber,
		CVC:        input.CVC,
		CardExp:    input.CardExp,
	}

	order, err := h.services.Purchase.CreateOrder(c.Request.Context(), order)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.JSON(http.StatusCreated, order)
}
