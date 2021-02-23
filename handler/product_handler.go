package handler

import (
	"fmt"
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

	products, _ := p.ProductRepo.SelectProducts(context.Request().Context())
	fmt.Println(len(products))
	for _, p := range products {
		fmt.Println(p.AttributesDb)
	}


	productId, err := uuid.NewUUID()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	productReq.ProductId = productId.String()

	_, err = p.ProductRepo.SaveProduct(context.Request().Context(), productReq)
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
	productReq := model.Product{}
	if err := context.Bind(&productReq); err != nil {
		log.Error(err.Error())
		return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err := p.ProductRepo.UpdateProduct(context.Request().Context(), productReq)
	if err != nil {
		return context.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Cập nhật sản phẩm thành công",
		Data:       nil,
	})
}

func (p ProductHandler) HandleDeleteAttrById(context echo.Context) error {
	attrId := context.Param("id")

	if len(attrId) == 0 {
		return context.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Thiếu thuộc tính id",
			Data:       nil,
		})
	}

	err := p.ProductRepo.DeleteProductAttr(context.Request().Context(), attrId)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xóa thuộc tính thành công",
		Data:       nil,
	})
}

func (p ProductHandler) HandleProductList(context echo.Context) error {
	mp := make(map[string]model.Product)
	products, _ := p.ProductRepo.SelectProducts(context.Request().Context())
	for _, p := range products {
		if pInMap, ok := mp[p.ProductId]; ok {
			pInMap.Attributes = append(pInMap.Attributes, p.AttributesDb)
			mp[p.ProductId] = pInMap
		} else {
			p.Attributes = append(p.Attributes, p.AttributesDb)
			mp[p.ProductId] = p
		}
	}

	var productsRes []model.Product
	for _, p := range mp {
		productsRes = append(productsRes, p)
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       productsRes,
	})
}
