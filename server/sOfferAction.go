package server

import (
	"net/http"

	"github.com/alf248/gotrade/database"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func offerAction(c echo.Context) error {

	user, authStatus, err := authenticate(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	actionForm := new(database.ActionForm)
	if err := c.Bind(actionForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	status2, err := database.OfferAction(objectId, user, actionForm)
	if err != nil {
		return c.JSON(status2, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
