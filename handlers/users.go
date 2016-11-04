package handlers

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-users/models"
	"net/http"
	"strconv"
)

func UsersCreate(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := user.Create().Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func UsersShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := user.FindById(id).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
