package model

type SignIn struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
