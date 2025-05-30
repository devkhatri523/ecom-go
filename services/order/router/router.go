package router

import (
	"github.com/devkhatri523/order-service/controller"
	"github.com/gin-gonic/gin"
)

func OrderRouter(orderController *controller.OrderController, orderLineController *controller.OrderLineController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/api/orders")
	router.POST("", orderController.CreateOrder)
	router.GET("/:orderId", orderController.FindByOrderId)
	router.GET("", orderController.FindAllOrder)
	router.GET("/orderLines/:orderLineId", orderLineController.FindAllByOrderId)
	return service
}
