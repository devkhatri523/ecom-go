package controller

import (
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	"github.com/devkhatri523/order-service/data/response"
	"github.com/devkhatri523/order-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderLineController struct {
	orderLineService service.OrderLineService
}

func NewOrderLineController(orderLineService service.OrderLineService) *OrderLineController {
	return &OrderLineController{
		orderLineService: orderLineService,
	}
}

func (ol *OrderLineController) FindAllByOrderId(ctx *gin.Context) {
	id := ctx.Param("orderId")
	orderId, err := strconv.Atoi(id)

	data, err := ol.orderLineService.FindAllByOrderId(int32(orderId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: utils.NewError(strconv.Itoa(500), err).Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, data)

}
