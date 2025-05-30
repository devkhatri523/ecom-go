package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"v01/data/request"
	"v01/data/response"
	"v01/service"
)

type PaymentController struct {
	paymentService service.PaymentService
}

func NewProductController(paymentService service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

func (paymentController *PaymentController) Create(ctx *gin.Context) {
	req := request.PaymentRequest{}
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	product, err := paymentController.paymentService.CreatePayment(req)
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
		Data:   product,
	}

	ctx.JSON(http.StatusOK, res)
}
