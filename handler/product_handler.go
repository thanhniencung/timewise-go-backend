package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"timewise/banana"
	"timewise/model"
	"timewise/repository"
)

type ProductHandler struct {
	ProductRepo repository.ProductRepo
}

func (p ProductHandler) HandleAddProduct(context echo.Context) error {
	productReq := model.Product{}
	if err := context.Bind(&productReq); err != nil {
		log.Error(err.Error())
		return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	productId, _ := uuid.NewUUID()
	productReq.ProductId = productId.String()

	_, err := p.ProductRepo.SaveProduct(context.Request().Context(), productReq)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = p.ProductRepo.AddProductAttribute(context.Request().Context(),
		productReq.ProductId, productReq.CollectionId, productReq.Attributes)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

func (p ProductHandler) HandleProductDetail(context echo.Context) error {
	productId := context.Param("id")

	product, err := p.ProductRepo.SelectProductById(context.Request().Context(), productId)
	if err != nil {
		if err == banana.ProductNotFound {
			return context.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       product,
	})
}

func (p ProductHandler) HandleEditProduct(context echo.Context) error {
	return nil
}

func (p ProductHandler) HandleProductList(context echo.Context) error {
	return nil
}
