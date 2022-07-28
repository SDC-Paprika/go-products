package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	products := router.Group("/products")

	products.GET("/", getProducts)
	products.GET("/:productId", getDetails)
	products.GET("/:productId/styles", getStyles)
	products.GET("/:productId/related", getRelated)

	router.Run(":6868")
}

/* TODO: Move handler functions to controllers */

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []string{"butts", "nuts"})
}

func getDetails(c *gin.Context) {
	productId := c.Param("productId")
	c.JSON(http.StatusOK, "Product ID: "+productId)
}

func getStyles(c *gin.Context) {
	productId := c.Param("productId")
	c.JSONP(http.StatusOK, "Product ID: "+productId)
}

func getRelated(c *gin.Context) {
	productId := c.Param("productId")
	c.IndentedJSON(http.StatusOK, "Product ID: "+productId)
}
