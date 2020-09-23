package model

import "time"

type User struct {
	UserID    string    `json:"-" db:"user_id, omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name, omitempty"`
	Email     string    `json:"email,omitempty" db:"email, omitempty"`
	Password  string    `json:"-" db:"password, omitempty"`
	Phone     string    `json:"phone,omitempty" db:"phone, omitempty"`
	Avatar    string    `json:"avatar,omitempty" db:"avatar, omitempty"`
	Address   *string   `json:"address,omitempty" db:"address, omitempty"`
	Role      string    `json:"-" db:"role, omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"createdAt" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
}
