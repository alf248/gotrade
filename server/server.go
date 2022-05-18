package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(port string, corsOrigin string) {

	e := echo.New()

	// Limit size of requests
	// https://echo.labstack.com/middleware/body-limit/
	// "The body limit is determined based on both Content-Length request header and actual content read, which makes it super secure." -echo docs
	// Limit can be specified as 4x or 4xB, where x is one of the multiple from K, M, G, T or P.
	e.Use(middleware.BodyLimit("1M"))

	// Limit rate of requests
	// https://echo.labstack.com/middleware/rate-limiter/
	// "By default an in-memory store is used for keeping track of requests" -echo docs
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// https://echo.labstack.com/middleware/cors/
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// Just a path to test that the server is running from a browser
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "OK") })

	// USERS
	e.POST("/users/:name", getUser)
	e.POST("/users/:name/edit", editUser)

	// LOGIN
	e.POST("/login", func(c echo.Context) error { return login(c) })
	e.POST("/logout", func(c echo.Context) error { return logout(c) })

	// OFFERS
	e.POST("/offers/new", func(c echo.Context) error { return newOffer(c) })
	e.POST("/offers", func(c echo.Context) error { return getOffers(c) })
	e.POST("/offers/:id", func(c echo.Context) error { return getOffer(c) })
	e.POST("/offers/:id/edit", func(c echo.Context) error { return editOffer(c) })
	e.POST("/offers/:id/action", func(c echo.Context) error { return offerAction(c) })

	// ORDERS
	e.POST("/orders/new", newVectors)
	e.POST("/orders", getVectors)
	e.POST("/orders/:id/action", vectorAction)

	// Start the server
	// this "fatal" logger will crash this program if the server fails somehow
	e.Logger.Fatal(e.Start(":" + port))

}
