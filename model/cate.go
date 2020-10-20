package model

import "time"

type Cate struct {
	CateId    string    `json:"cateId,omitempty" db:"cate_id, omitempty"`
	CateName  string    `json:"cateName,omitempty" db:"cate_name, omitempty"`
	CateImage string    `json:"cateImage,omitempty" db:"cate_image, omitempty"`
	CreatedAt time.Time `json:"createdAt" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at, omitempty"`
}
