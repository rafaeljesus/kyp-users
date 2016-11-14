package api

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-users/models"
	"net/http"
	"strconv"
)

const UNAUTHORIZED = "Unauthorized"

func (env *Env) UsersCreate(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := env.Repo.CreateUser(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (env *Env) UsersShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	if err := env.Repo.FindUserById(&user, id).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (env *Env) UsersAuthenticate(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	user := models.User{}
	if err := env.Repo.FindUserByEmail(&user, u.Email).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	if invalid, _ := user.VerifyPassword(u.Password); !invalid {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	return c.JSON(http.StatusOK, user)
}
