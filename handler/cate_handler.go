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

type CateHandler struct {
	CateRepo repository.CateRepo
}

func (c CateHandler) HandleAddCate(context echo.Context) error {
	req := model.Cate{}
	if err := context.Bind(&req); err != nil {
		log.Error(err.Error())
		return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	cateId, _ := uuid.NewUUID()
	req.CateId = cateId.String()

	fmt.Println(req)

	cate, err := c.CateRepo.SaveCate(context.Request().Context(), req)
	if err != nil {
		return context.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       cate,
	})
}

func (c CateHandler) HandleEditCate(context echo.Context) error {
	req := model.Cate{}
	if err := context.Bind(&req); err != nil {
		log.Error(err.Error())
		return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err := c.CateRepo.UpdateCate(context.Request().Context(), req)
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

func (c CateHandler) HandleCateDetail(context echo.Context) error {
	cateId := context.Param("id")

	cate, err := c.CateRepo.SelectCateById(context.Request().Context(), cateId)
	if err != nil {
		if err == banana.CateNotFound {
			return context.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return context.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       cate,
	})
}

func (c CateHandler) HandleCateList(context echo.Context) error {
	cates, err := c.CateRepo.SelectCates(context.Request().Context())
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
		Data:       cates,
	})
}
