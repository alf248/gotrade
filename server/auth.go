package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"

	"github.com/labstack/echo/v4"
)

var fireapp *firebase.App

func authenticate_through_firebase(c echo.Context) (user *forms.User, httpStatus int, e error) {

	idToken, err := getToken(c)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	firebaseId, status, err := verify(idToken)
	if err != nil {
		return nil, status, err
	}

	user, status, err = database.GetUserByFirebaseId(firebaseId)
	if err != nil {

		if status == http.StatusNotFound {

			user := forms.NewUser{FID: firebaseId}

			_, err := database.NewUser(user)
			if err != nil {
				return nil, http.StatusInternalServerError, err
			}

		} else {
			return nil, status, err
		}

	}

	return user, http.StatusOK, nil
}

func verify(idToken string) (firebaseUserId string, httpStatus int, e error) {

	//fmt.Printf("Authenticating: got token: %+v", idToken)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := fireapp.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	//log.Printf("Verified ID token: %v\n", token)

	return token.UID, http.StatusOK, nil
}

func getToken(c echo.Context) (string, error) {

	errMsg := "Authorization header malformed: should be 'Bearer <token>'"

	for key, values := range c.Request().Header {

		if key == "Authorization" {

			//fmt.Println("Found authorization key:", key, "values:", values)

			if len(values) > 0 {

				s := strings.Fields(values[0])

				if len(s) > 1 {
					//println("token:", s[1])
					return s[1], nil
				} else {
					return "", errors.New(errMsg + " got: " + values[0])
				}

			}

			return "", errors.New(errMsg + " got: " + values[0])
			/*
				for _, value := range values {
					fmt.Println("val:", value)
				}
			*/

		}

	}

	return "", errors.New("Found no Authorization header in http request")
}
