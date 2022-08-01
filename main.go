package main

import (
	"github.com/SDC-Paprika/go-products/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	products := router.Group("/products")
	{
		product := new(controllers.ProductsController)
		details := new(controllers.DetailsController)
		related := new(controllers.RelatedController)
		styles := new(controllers.StylesController)

		products.GET("/", product.GetProducts)
		products.GET("/:productId", details.GetDetails)
		products.GET("/:productId/styles", styles.GetStyles)
		products.GET("/:productId/related", related.GetRelated)
	}

	router.Run(":6868")
}
