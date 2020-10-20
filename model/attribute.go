package model

import "time"

type Attribute struct {
	AttrId       string    `json:"attrId,omitempty" db:"attr_id, omitempty"`
	CollectionId string    `json:"collectionId,omitempty" db:"collection_id, omitempty"`
	ProductId    string    `json:"productId,omitempty" db:"product_id, omitempty"`
	AttrName     string    `json:"attrName,omitempty" db:"attr_name, omitempty"`
	Size         int    	`json:"size,omitempty" db:"size, omitempty"`
	Price        float64    `json:"price,omitempty" db:"price, omitempty"`
	Promotion    float64   `json:"promotion,omitempty" db:"promotion, omitempty"`
	Quantity     int    `json:"quantity,omitempty" db:"quantity, omitempty"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at, omitempty"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at, omitempty"`
}
