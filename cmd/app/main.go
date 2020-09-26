package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"timewise/db"
	"timewise/handler"
	"timewise/model"
	"timewise/repository"
	"timewise/router"
)

func main() {
	var cfg model.Config
	loadConfig(&cfg)
	setupEnv(&cfg)

	var sql = new(db.SQL)
	sql.Connect(cfg)
	defer sql.Close()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	userHandler := handler.UserHandler{
		UserRepo: repository.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}

func setupEnv(cfg *model.Config) {
	jwtExpires := strconv.Itoa(cfg.Server.JwtExpires)
	os.Setenv("JwtExpires", jwtExpires)
}

func loadConfig(cfg *model.Config) {
	f, err := os.Open("../../env.dev.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}
