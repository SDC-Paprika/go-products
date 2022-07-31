package controllers

import (
	"net/http"
	"strconv"

	"github.com/SDC-Paprika/go-products/models"
	"github.com/gin-gonic/gin"
)

type RelatedController struct{}

var relatedModel models.RelatedModel

func (ctrl RelatedController) GetRelated(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Bad Product ID: " + c.Param("productId")})
	}

	results, err := relatedModel.Get(productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get related products"})
	}

	c.IndentedJSON(http.StatusOK, results)
}
