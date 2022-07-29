package controllers

import (
	"net/http"

	"github.com/SDC-Paprika/go-products/models"
	"github.com/gin-gonic/gin"
)

type ProductsController struct{}

var productsModel = new(models.ProductsModel)

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []string{"butts", "nuts"})
}
