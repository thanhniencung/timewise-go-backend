package router

import (
	"github.com/labstack/echo/v4"
	"timewise/handler"
	"timewise/middleware"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	// user
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile", api.UserHandler.HandleProfile, middleware.JWTMiddleware())
}
