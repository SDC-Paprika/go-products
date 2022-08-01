package controllers

import (
	"net/http"
	"strconv"

	"github.com/SDC-Paprika/go-products/models"
	"github.com/gin-gonic/gin"
)

type StylesController struct{}

var stylesModel = new(models.StylesModel)

func (ctrl StylesController) GetStyles(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "error while parsing product id: " + err.Error()})
	}
	if productId < 1 || productId > 1000011 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "bad product id: " + c.Param("productId")})
	}

	results, err := stylesModel.Get(productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "could not fetch styles"})
	}
	c.JSON(http.StatusOK, results)
}
