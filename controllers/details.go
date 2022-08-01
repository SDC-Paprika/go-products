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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "error parsing product id: " + err.Error()})
	}
	if productId < 1 || productId > 1000011 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "bad product id: " + c.Param("productId")})
	}

	results, err := detailsModel.Get(productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Could not get details"})
		return
	}

	c.JSON(http.StatusOK, results)
}
