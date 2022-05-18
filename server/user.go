package server

import (
	"net/http"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {

	name := c.Param("name")

	token, _, err := getTokenAndEmail(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, status, err := database.GetUser(name)
	if user == nil {
		return c.JSON(status, err.Error())
	}

	if token != user.Token {
		user.MakePrivate()
	}

	return c.JSON(http.StatusOK, user)
}

func editUser(c echo.Context) error {

	editForm := new(forms.EditUser)
	if err := c.Bind(editForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if editForm.Phone == "" {
		editForm.Phone = " "
	}

	user, authStatus, err := authenticate(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	user, status, err := database.EditUser(user.Name, editForm)
	if user == nil {
		return c.JSON(int(status), err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
