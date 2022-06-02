package server

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var path_to_service_account_key = "ServiceAccountKey.json"

const firebaseProjectId = "trade-51694"

func StartStandardServer(port string, corsOrigin string) {

	_, err := ioutil.ReadFile(path_to_service_account_key)
	if err != nil {
		path_to_service_account_key = "../" + path_to_service_account_key
		_, err = ioutil.ReadFile(path_to_service_account_key)
		if err != nil {
			log.Fatalf("error reading service account key")
		}
	}

	opt := option.WithCredentialsFile(path_to_service_account_key)
	config := &firebase.Config{ProjectID: firebaseProjectId}
	fireapp, err = firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	e := echo.New()

	// Limit size of requests
	// https://echo.labstack.com/middleware/body-limit/
	// Limit can be specified as 4x or 4xB, where x is one of the multiple from K, M, G, T or P.
	e.Use(middleware.BodyLimit("1M"))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "gotrade server is running") })

	// OFFERS
	e.POST("/offers/new", newOffer)
	e.POST("/offers", getOffers)
	e.POST("/offers/:id", getOffer)
	e.POST("/offers/:id/edit", editOffer)
	e.POST("/offers/:id/action", offerAction)

	// ORDERS
	e.POST("/orders/new", newOrders)
	e.POST("/orders", getOrders)
	e.POST("/orders/:id/action", orderAction)

	// Start the server
	e.Logger.Fatal(e.Start(":" + port))
}
