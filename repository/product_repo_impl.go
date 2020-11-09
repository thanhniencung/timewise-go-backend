package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"time"
	"timewise/banana"
	"timewise/db"
	"timewise/model"
)

type ProductRepoImpl struct {
	sql *db.SQL
}

// NewUserRepo create object working with user logic
func NewProductRepo(sql *db.SQL) ProductRepo {
	return ProductRepoImpl{
		sql: sql,
	}
}

func (p ProductRepoImpl) SaveProduct(context context.Context, product model.Product) (model.Product, error) {
	statement := `
		INSERT INTO products(
				product_id, product_name, product_image, product_des, 
				cate_id, collection_id, created_at, updated_at)
		VALUES(:product_id, :product_name, :product_image, :product_des, 
				:cate_id, :collection_id, :created_at, :updated_at)
	`
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	_, err := p.sql.Db.NamedExecContext(context, statement, product)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return product, errors.New("Sản phẩm này đã tồn tại")
			}
		}
		return product, errors.New("Tạo Sản phẩm thất bại")
	}

	return product, nil
}

func (p ProductRepoImpl) AddProductAttribute(context context.Context,
	productId string, collectionId string, attributes []model.Attribute) error {

	statement := `
		INSERT INTO attributes(
				attr_id, p_id, col_id, attr_name, size, 
				price, promotion, quantity, created_at, updated_at)
		VALUES(:attr_id, :p_id, :col_id, :attr_name, 
				:size, :price, :promotion, :quantity, :created_at, :updated_at)
	`

	tx := p.sql.Db.MustBegin()
	for _, attr := range attributes {
		uuid, _ := uuid.NewUUID()
		attr.AttrId = uuid.String()
		attr.ProductId = productId
		attr.CollectionId = collectionId

		now := time.Now()
		attr.CreatedAt = now
		attr.UpdatedAt = now

		_, err := tx.NamedExecContext(context, statement, attr)
		if err != nil {
			tx.Rollback()
			log.Error(err.Error())
			if err, ok := err.(*pq.Error); ok {
				if err.Code.Name() == "unique_violation" {
					return errors.New("Có thuộc tính này đã tồn tại")
				}
			}
			return errors.New("Thêm thuộc tính thất bại")
		}
	}
	tx.Commit()

	return nil
}

func (p ProductRepoImpl) SelectProductById(context context.Context, productId string) (model.Product, error) {
	var product = model.Product{}
	var attrs []model.Attribute

	statement := `SELECT * FROM products WHERE product_id=$1`
	err := p.sql.Db.GetContext(context, &product, statement, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, banana.ProductNotFound
		}
		log.Error(err.Error())
		return product, err
	}

	statement = `SELECT * FROM attributes WHERE product_id=$1`
	err = p.sql.Db.SelectContext(context, &attrs, statement, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, errors.New("Sản phẩm này không tồn tại thuộc tính nào")
		}
		log.Error(err.Error())
		return product, err
	}

	product.Attributes = attrs

	return product, nil
}

func (p ProductRepoImpl) UpdateProduct(context context.Context, product model.Product) error {
	return nil
}

func (p ProductRepoImpl) SelectProducts(context context.Context) ([]model.Product, error) {
	var products []model.Product
	sql := `SELECT
	      products.*,
	      attributes.attr_id,
	      attributes.col_id,
	      attributes.attr_name,
	      attributes.size,
	      attributes.price,
	      attributes.promotion,
	      attributes.quantity
	    FROM
	      products JOIN attributes 
	      ON products.product_id = attributes.p_id;`
	err := p.sql.Db.Select(&products, sql)
	return products, err
}




