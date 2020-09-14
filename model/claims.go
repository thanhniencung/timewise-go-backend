package model

import "github.com/dgrijalva/jwt-go"

// JwtCustomClaims
type JwtCustomClaims struct {
	UserId string
	Role   string
	jwt.StandardClaims
}
