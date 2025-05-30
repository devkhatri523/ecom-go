package router

import (
	"github.com/gin-gonic/gin"
	"v01/controller"
)

func PaymentRouter(paymentController *controller.PaymentController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/api/payment")
	router.POST("", paymentController.Create)
	return service
}
