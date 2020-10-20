package model

import "time"

type Collection struct {
	CateId         string    `json:"cateId,omitempty" db:"cate_id, omitempty"`
	CollectionId   string    `json:"collectionId,omitempty" db:"collection_id, omitempty"`
	CollectionName string    `json:"collectionName,omitempty" db:"collection_name, omitempty"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at, omitempty"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at, omitempty"`
}
