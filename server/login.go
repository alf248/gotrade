package server

import (
	"net/http"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"

	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {

	loginForm := new(forms.LoginForm)
	if err := c.Bind(loginForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := database.Login(loginForm.Email, loginForm.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	return c.JSON(http.StatusOK, user)
}

func logout(c echo.Context) error {

	logoutForm := new(forms.LogoutForm)
	if err := c.Bind(logoutForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := database.Logout(logoutForm.Email, logoutForm.Token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	return c.JSON(http.StatusOK, "")
}
