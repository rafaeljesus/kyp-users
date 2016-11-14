package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-users/api"
	"github.com/rafaeljesus/kyp-users/models"
	"log"
	"os"
)

var KYP_USERS_PORT = os.Getenv("KYP_USERS_PORT")
var KYP_USERS_DB = os.Getenv("KYP_USERS_DB")

func main() {
	db, err := models.NewDB(KYP_USERS_DB)
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.User{})

	env := &api.Env{db}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/v1")
	v1.GET("/healthz", env.HealthzIndex)
	v1.POST("/users", env.UsersCreate)
	v1.GET("/users/:id", env.UsersShow)
	v1.POST("/users/authenticate", env.UsersAuthenticate)

	log.Print("Starting Kyp Users Service at port ", KYP_USERS_PORT)

	e.Run(fasthttp.New(":" + KYP_USERS_PORT))
}
