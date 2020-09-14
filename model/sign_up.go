package model

type SignUp struct {
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
