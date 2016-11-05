package handlers

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-users/models"
	"net/http"
	"strconv"
)

const UNAUTHORIZED = "Unauthorized"

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

func UsersAuthenticate(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	user := models.User{}
	if err := user.FindByEmail(u.Email).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	if invalid, _ := user.VerifyPassword(u.Password); !invalid {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	return c.JSON(http.StatusOK, user)
}
