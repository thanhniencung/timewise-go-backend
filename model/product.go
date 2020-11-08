package model

import "time"

type Product struct {
	Attributes    []Attribute `json:"attributes,omitempty"`
	ProductId    string    `json:"productId,omitempty" db:"product_id, omitempty"`
	ProductName  string    `json:"productName,omitempty" db:"product_name, omitempty"`
	ProductImage string    `json:"productImage,omitempty" db:"product_image, omitempty"`
	CateId       string    `json:"cateId,omitempty" db:"cate_id, omitempty"`
	Description  string    `json:"description,omitempty" db:"product_des, omitempty"`
	CollectionId string    `json:"collectionId,omitempty" db:"collection_id, omitempty"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at, omitempty"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at, omitempty"`
}