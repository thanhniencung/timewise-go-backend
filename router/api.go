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
	user.GET("/list", api.UserHandler.HandleListUsers, middleware.JWTMiddleware())
}

func (api *API) SetupAdminRouter() {
	//admin
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.UserHandler.GenToken)
	admin.POST("/sign-in", api.UserHandler.HandleAdminSignIn)
	admin.POST("/sign-up", api.UserHandler.HandleAdminSignUp, middleware.JWTMiddleware())
}
