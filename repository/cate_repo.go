package repository

import (
	"context"
	"timewise/model"
)

type CateRepo interface {
	SaveCate(context context.Context, cate model.Cate) (model.Cate, error)
	DeleteCate(context context.Context, cateId string) error
	UpdateCate(context context.Context, cate model.Cate) error
	SelectCateById(context context.Context, cateId string) (model.Cate, error)
	SelectCates(context context.Context) ([]model.Cate, error)
}
