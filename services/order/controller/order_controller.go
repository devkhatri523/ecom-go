package controller

import (
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	"github.com/devkhatri523/order-service/data/request"
	"github.com/devkhatri523/order-service/data/response"
	"github.com/devkhatri523/order-service/mapper"
	"github.com/devkhatri523/order-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (o *OrderController) FindAllOrder(ctx *gin.Context) {
	req := request.GetAllOrderRequest{}
	err := ctx.ShouldBindJSON(req)
	filter := mapper.MapToDomainFilter(req.Filters)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    404,
			Message: utils.NewError(strconv.Itoa(500), err).Error(),
		})
		return
	}
	data, _, err := o.orderService.FindAllOrder(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: utils.NewError(strconv.Itoa(500), err).Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, data)

}
func (o *OrderController) CreateOrder(ctx *gin.Context) {
	req := request.OrderRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    404,
			Message: utils.NewError(strconv.Itoa(500), err).Error(),
		})
		return
	}
	order, err := o.orderService.CreateOrder(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    404,
			Message: utils.NewError(strconv.Itoa(500), err).Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, order)

}

func (o *OrderController) FindByOrderId(ctx *gin.Context) {
	id := ctx.Param("orderId")
	orderId, err := strconv.Atoi(id)

	data, err := o.orderService.FindByOrderId(int32(orderId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: utils.NewError(strconv.Itoa(500), err).Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, data)

}
