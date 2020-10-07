package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"timewise/model"
	"timewise/security"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.JwtKey),
	}

	return middleware.JWTWithConfig(config)
}

func CheckAdminRole() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenData := c.Get("user").(*jwt.Token)
			claims := tokenData.Claims.(*model.JwtCustomClaims)

			if claims.Role != model.ADMIN.String() {
				return c.JSON(http.StatusForbidden, model.Response{
					StatusCode: http.StatusForbidden,
					Message:    "Không cho phép truy cập",
					Data:       nil,
				})
			}

			return next(c)
		}
	}
}
