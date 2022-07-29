package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRelated(c *gin.Context) {
	productId := c.Param("productId")
	c.IndentedJSON(http.StatusOK, "Product ID: "+productId)
}
