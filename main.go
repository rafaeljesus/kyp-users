package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-users/db"
	"github.com/rafaeljesus/kyp-users/handlers"
	"github.com/rafaeljesus/kyp-users/models"
	"log"
	"os"
)

var KYP_USERS_PORT = os.Getenv("KYP_USERS_PORT")

func main() {
	db.Connect()
	db.Repo.AutoMigrate(&models.User{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/v1")
	v1.GET("/healthz", handlers.HealthzIndex)
	v1.POST("/users", handlers.UsersCreate)
	v1.GET("/users/:id", handlers.UsersShow)
	v1.POST("/users/authenticate", handlers.UsersAuthenticate)

	log.Print("Starting Kyp Users Service at port ", KYP_USERS_PORT)

	e.Run(fasthttp.New(":" + KYP_USERS_PORT))
}
