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

	//admin
	admin := api.Echo.Group("/admin")
	admin.POST("/sign-in", api.UserHandler.HandleAdminSignIn)
	admin.GET("/token", api.UserHandler.GenToken)
	admin.GET("/sign-up", api.UserHandler.HandleAdminSignUp,
		middleware.JWTMiddleware())
}
