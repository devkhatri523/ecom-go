package router

import (
	"github.com/gin-gonic/gin"
	"v01/controller"
)

func ProductRouter(productController *controller.ProductController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/products")

	router.GET("", productController.FindAll)
	router.GET("/:id", productController.FindById)
	router.POST("", productController.Create)
	router.PATCH("/:id", productController.Update)
	router.DELETE("/:id", productController.Delete)
	router.POST("/purchase", productController.PurchaseProducts)
	return service
}
