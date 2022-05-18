package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"

	"github.com/labstack/echo/v4"
)

func newVectors(c echo.Context) error {

	user, authStatus, err := authenticate(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	if user.VectorsMade >= database.MAX_VECTORS_PER_USER {
		return c.JSON(http.StatusForbidden, errors.New("You have made too many vectors"))
	}

	idList := new([]string)
	if err := c.Bind(idList); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	offers, status, err := database.GetOffersById(idList)
	if err != nil {
		return c.JSON(status, err.Error())
	}

	var vectors []forms.VectorInOut
	// Iterate the offers and do some verification:
	for _, gun := range offers {
		v := forms.VectorInOut{}
		v.Init(gun, user.Name)
		err := v.AddUserAsHeadOrTail(user.Name)
		if err != nil {
			return c.String(http.StatusConflict, err.Error())
		}
		vectors = append(vectors, v)
	}

	status, err = database.NewVectors(vectors)
	if err != nil {
		return c.String(status, err.Error())
	}

	// Update user stats
	var editUserForm = new(forms.EditUser)
	editUserForm.VectorsMade = user.VectorsMade + len(vectors)
	database.EditUser(user.Name, editUserForm)

	return c.JSON(http.StatusOK, "ok")
}

func getVectors(c echo.Context) error {

	searchForm := new(forms.VectorSearch)
	if err := c.Bind(searchForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := searchForm.Curate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	vectors, status, err := database.GetVectors(searchForm.SortUp, searchForm.Active)
	if err != nil {
		return c.JSON(int(status), err.Error())
	}

	// use this for large json:
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(vectors)
}

func vectorAction(c echo.Context) error {

	user, authStatus, err := authenticate(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	actionForm := new(forms.VectorAction)
	if err := c.Bind(actionForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	switch actionForm.Action {
	case "remove":
		return vectorActionRemove(c, user)
	default:
		return c.JSON(http.StatusBadRequest, "action must be remove or pause. got "+actionForm.Action)
	}

}

func vectorActionRemove(c echo.Context, user *forms.User) error {

	orderId := c.Param("id")

	status, err := database.DeleteVector(orderId, user.Name)
	if err != nil {
		return c.JSON(int(status), err.Error())
	}

	return c.JSON(http.StatusOK, "ok")
}
