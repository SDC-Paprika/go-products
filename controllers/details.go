package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getDetails(c *gin.Context) {
	productId := c.Param("productId")
	c.JSON(http.StatusOK, "Product ID: "+productId)
}
