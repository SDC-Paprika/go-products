package controllers

import (
	"net/http"
	"strconv"

	"github.com/SDC-Paprika/go-products/models"
	"github.com/gin-gonic/gin"
)

type DetailsController struct{}

var detailsModel = new(models.DetailsModel)

func (ctrl DetailsController) GetDetails(c *gin.Context) {
	var productId int
	var err error

	if productId, err = strconv.Atoi(c.Param("productId")); err != nil {
		productId = 1
	}

	results, err := detailsModel.Get(productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get details"})
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}
