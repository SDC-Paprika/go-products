package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStyles(c *gin.Context) {
	productId := c.Param("productId")
	c.JSONP(http.StatusOK, "Product ID: "+productId)
}
