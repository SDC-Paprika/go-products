package controllers

import (
	"net/http"
	"strconv"

	"github.com/SDC-Paprika/go-products/models"
	"github.com/gin-gonic/gin"
)

type ProductsController struct{}

var productsModel = new(models.ProductsModel)

func (ctrl ProductsController) GetProducts(c *gin.Context) {
	var page, count int
	var err error

	if page, err = strconv.Atoi(c.DefaultQuery("page", "1")); err != nil {
		page = 1
	}
	if count, err = strconv.Atoi(c.DefaultQuery("count", "5")); err != nil {
		count = 5
	}

	results, err := productsModel.Get(page, count)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get products"})
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}
