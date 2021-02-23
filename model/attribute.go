package model

import "time"

type Attribute struct {
	AttrId       string    `json:"attrId,omitempty" db:"attr_id, omitempty"`
	CollectionId string    `json:"-" db:"col_id, omitempty"`
	ProductId    string    `json:"productId,omitempty" db:"p_id, omitempty"`
	AttrName     string    `json:"attrName,omitempty" db:"attr_name, omitempty"`
	Size         int       `json:"size,omitempty" db:"size, omitempty"`
	Price        float64   `json:"price,omitempty" db:"price, omitempty"`
	Promotion    float64   `json:"promotion,omitempty" db:"promotion, omitempty"`
	Quantity     int       `json:"quantity,omitempty" db:"quantity, omitempty"`
	CreatedAt    time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt    time.Time `json:"-" db:"updated_at, omitempty"`
}
