package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"

	"github.com/labstack/echo/v4"
)

func newOrders(c echo.Context) error {

	user, authStatus, err := authenticate_through_firebase(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	if user.OrdersMade >= database.MAX_VECTORS_PER_USER {
		return c.JSON(http.StatusForbidden, errors.New("you have made too many orders"))
	}

	idList := new([]string)
	if err := c.Bind(idList); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	offers, status, err := database.GetOffersById(idList)
	if err != nil {
		return c.JSON(status, err.Error())
	}

	var orders []forms.Order
	for _, offer := range offers {
		order := forms.Order{}
		err = order.Init(offer, user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		orders = append(orders, order)
	}

	status, err = database.NewOrders(orders)
	if err != nil {
		return c.String(status, err.Error())
	}

	// Update user stats
	var editUserForm = new(forms.EditUser)
	editUserForm.VectorsMade = user.OrdersMade + len(orders)
	database.EditUser(user.Name, editUserForm)

	return c.JSON(http.StatusOK, "ok")
}

func getOrders(c echo.Context) error {

	user, authStatus, err := authenticate_through_firebase(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	form := new(forms.SearchOrders)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = form.Curate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	println("")
	println("")
	fmt.Printf("got form %+v", form)

	orders, status, err := database.GetOrders(form.SortUp, form.Active, form.AsGiver, user.FID)
	if err != nil {
		return c.JSON(int(status), err.Error())
	}

	// use this for large json:
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(orders)
}

func orderAction(c echo.Context) error {

	user, authStatus, err := authenticate_through_firebase(c)
	if err != nil {
		return c.JSON(authStatus, err.Error())
	}

	actionForm := new(forms.OrderAction)
	if err := c.Bind(actionForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	switch actionForm.Action {
	case "remove":
		return cancelOrder(c, user)
	default:
		return c.JSON(http.StatusBadRequest, "action must be remove or pause. got "+actionForm.Action)
	}

}

func cancelOrder(c echo.Context, user *forms.User) error {

	orderId := c.Param("id")

	status, err := database.DeleteOrder(orderId, user.FID)
	if err != nil {
		return c.JSON(int(status), err.Error())
	}

	return c.JSON(http.StatusOK, "ok")
}
