package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"v01/data/request"
	"v01/data/response"
	"v01/service"
	"v01/validator"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (productController *ProductController) FindAll(ctx *gin.Context) {
	data, err := productController.ProductService.FindAll()
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

func (productController *ProductController) Create(ctx *gin.Context) {
	req := request.ProductRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	err = validator.ValidateProductCreateRequest(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	product, err := productController.ProductService.CreateProduct(req)
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

func (productController *ProductController) FindById(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	data, err := productController.ProductService.FindById(int32(id))
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

func (productController *ProductController) PurchaseProducts(ctx *gin.Context) {
	var req []request.ProductPurchaseRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	purchaseProducts, err := productController.ProductService.PurchaseProducts(req)
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
		Data:   purchaseProducts,
	}

	ctx.JSON(http.StatusOK, res)

}

func (productController *ProductController) Update(ctx *gin.Context) {
	req := request.UpdateProductRequest{}
	err := ctx.ShouldBindJSON(&req)

	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)

	_, err = productController.ProductService.FindById(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}
	req.Id = int32(id)
	err = productController.ProductService.Update(req)
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
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}

func (productController *ProductController) Delete(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)

	_, err = productController.ProductService.FindById(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	Id, err := productController.ProductService.Delete(int32(id))
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
		Data:   Id,
	}

	ctx.JSON(http.StatusOK, res)
}
