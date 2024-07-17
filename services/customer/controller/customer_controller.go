package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"v01/data/request"
	"v01/data/response"
	"v01/service"
	"v01/validator"
)

type CustomerController struct {
	customerService service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: service,
	}
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	req := request.CustomerRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	err = validator.ValidateCustomerCreateRequest(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	customer, err := c.customerService.CreateCustomer(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   customer,
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	customerId := ctx.Param("customerId")

	_, err := c.customerService.FindById(customerId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	customerId, err = c.customerService.DeleteCustomer(customerId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   fmt.Sprintf("customer with id:%s deleted sucessfully ", customerId),
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	req := request.CustomerUpdateRequest{}
	err := ctx.ShouldBindJSON(&req)

	customerId := ctx.Param("customerId")
	customerId, err = c.customerService.UpdateCustomer(customerId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   fmt.Sprintf("Customer with Id : %s updated sucessfully", customerId),
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *CustomerController) FindAllCustomerCustomer(ctx *gin.Context) {
	data, err := c.customerService.FindAllCustomer()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *CustomerController) FindCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("customerId")

	data, err := c.customerService.FindById(customerId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)

}

func (c *CustomerController) IsCustomerExists(ctx *gin.Context) {
	customerId := ctx.Param("customerId")

	data, err := c.customerService.IsCustomerExists(customerId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)

}
