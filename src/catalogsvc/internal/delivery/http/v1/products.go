package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("/", h.getAllProducts)
		products.GET("/:id", h.getProductByID)
	}
}

func (h *Handler) getAllProducts(c *gin.Context) {
	products, err := h.services.Products.GetAll(c.Request.Context())
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *Handler) getProductByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {

		h.newBadRequestResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.Products.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			h.newBadRequestResponse(c, http.StatusNotFound, err.Error())
			return
		}

		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}
