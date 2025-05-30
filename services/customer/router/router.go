package router

import (
	"github.com/gin-gonic/gin"
	"v01/controller"
)

func CustomerRouter(customerController *controller.CustomerController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/api/customers")

	router.GET("", customerController.FindAllCustomerCustomer)
	router.GET("/:customerId", customerController.FindCustomerById)
	router.POST("", customerController.CreateCustomer)
	router.PATCH("/:customerId", customerController.UpdateCustomer)
	router.DELETE("/:customerId", customerController.DeleteCustomer)
	router.GET("/customerexists/:customerId", customerController.IsCustomerExists)
	return service
}
