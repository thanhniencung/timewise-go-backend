package router

import (
	"github.com/labstack/echo/v4"
	"timewise/handler"
	"timewise/middleware"
)

type API struct {
	Echo           *echo.Echo
	UserHandler    handler.UserHandler
	CateHandler    handler.CateHandler
	ProductHandler handler.ProductHandler
}

func (api *API) SetupRouter() {
	// user
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile", api.UserHandler.HandleProfile, middleware.JWTMiddleware())
	user.GET("/list", api.UserHandler.HandleListUsers, middleware.JWTMiddleware())

	// categories
	categories := api.Echo.Group("/cate",
		middleware.JWTMiddleware(),
		middleware.CheckAdminRole())

	categories.POST("/add", api.CateHandler.HandleAddCate)
	categories.PUT("/edit", api.CateHandler.HandleEditCate)
	categories.GET("/detail/:id", api.CateHandler.HandleCateDetail)
	categories.GET("/list", api.CateHandler.HandleCateList)

	// product
	product := api.Echo.Group("/product",
		middleware.JWTMiddleware(),
		middleware.CheckAdminRole())

	product.POST("/add", api.ProductHandler.HandleAddProduct)
	product.DELETE("/attr/:id", api.ProductHandler.HandleDeleteAttrById)
	product.PUT("/edit", api.ProductHandler.HandleEditProduct)
	product.GET("/list", api.ProductHandler.HandleProductList)
	product.GET("/detail/:id", api.ProductHandler.HandleProductDetail)
}

func (api *API) SetupAdminRouter() {
	//admin
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.UserHandler.GenToken)
	admin.POST("/sign-in", api.UserHandler.HandleAdminSignIn)
	admin.POST("/sign-up", api.UserHandler.HandleAdminSignUp, middleware.JWTMiddleware())
}
