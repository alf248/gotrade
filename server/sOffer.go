package server

import (
	"encoding/json"
	"net/http"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo/v4"
)

func newOffer(c echo.Context) error {

	user, authStatus, err := authenticate(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	if user.OffersMade >= database.MAX_OFFERS_PER_USER {
		return c.JSON(http.StatusForbidden, "you have made too many offers")
	}

	form := new(forms.NewOfferForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = form.Curate(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newId, err := database.NewOffer(user, form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Update user stats
	var editUserForm = new(forms.EditUser)
	editUserForm.OffersMade = user.OffersMade + 1
	database.EditUser(user.Name, editUserForm)

	return c.JSON(http.StatusOK, newId)
}

func getOffer(c echo.Context) error {

	id := c.Param("id")

	vector, status, err := database.GetOffer(id)
	if err != nil {
		return c.JSON(status, err.Error())
	}

	return c.JSON(http.StatusOK, vector)
}

func getOffers(c echo.Context) error {

	searchForm := new(forms.SearchOffers)
	if err := c.Bind(searchForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := searchForm.Verify()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var vectors []forms.Offer

	vectors, status, err := database.SearchOffers(searchForm)
	if err != nil {
		return c.JSON(int(status), err.Error())
	}

	// use this for large json:
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(vectors)
}

func editOffer(c echo.Context) error {

	user, authStatus, err := authenticate(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	form := new(forms.NewOfferForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	status, err := database.EditOffer(objectId, user, form)
	if err != nil {
		return c.JSON(int(status), err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
