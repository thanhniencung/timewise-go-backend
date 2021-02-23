package repository

import (
	"context"
	"timewise/model"
)

type ProductRepo interface {
	SaveProduct(context context.Context, product model.Product) (model.Product, error)
	AddProductAttribute(context context.Context, productId string, collectionId string, attributes []model.Attribute) error
	UpdateProduct(context context.Context, product model.Product) error
	SelectProductById(context context.Context, productId string) (model.Product, error)
	SelectProducts(context context.Context) ([]model.Product, error)
	DeleteProductAttr(context context.Context, attrId string) error
}
