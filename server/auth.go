package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"

	"github.com/labstack/echo/v4"
)

func authenticate(c echo.Context) (*forms.User, int, error) {

	token, email, err := getTokenAndEmail(c)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	user, status, err := database.CheckUser(email, token)
	if err != nil {
		return nil, status, err
	}

	return user, status, nil
}

func getTokenAndEmail(c echo.Context) (string, string, error) {

	errMsg := "Authorization header malformed: should be 'Bearer <token> <email>'"

	for key, values := range c.Request().Header {

		if key == "Authorization" {

			//fmt.Println("Found authorization key:", key, "values:", values)

			if len(values) > 0 {

				s := strings.Fields(values[0])

				if len(s) > 2 {
					//println("token:", s[1])
					return s[1], s[2], nil
				} else {
					return "", "", errors.New(errMsg + " got: " + values[0])
				}

			}

			return "", "", errors.New(errMsg + " got: " + values[0])
			/*
				for _, value := range values {
					fmt.Println("val:", value)
				}
			*/

		}

	}

	return "", "", errors.New("Found no Authorization header in http request")
}
