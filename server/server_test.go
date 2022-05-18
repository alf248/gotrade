package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/database/mock"
	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const mongoURL = "mongodb://localhost:27017"

var userEmail = "joe@example.com"
var userPassword = "pass"
var sessionToken string

func TestMain(m *testing.M) {

	useTestDatabase()
	database.DropOffersCollection()
	mock.AddMockOffers(4)

	code := m.Run()

	os.Exit(code)
}

// Should fail because not logged in yet
func TestNewOfferTooEarly(t *testing.T) {

	c, rec := setup2("/offers/new", `{"name":"Flashy Car"}`)
	if assert.NoError(t, newOffer(c)) {
		assert.NotEqual(t, http.StatusOK, rec.Code)
	}
}

// This will test the login, and also login a user, which will be used in following tests
func TestLogin(t *testing.T) {

	body := `{"email":"` + userEmail + `", "password":"` + userPassword + `"}`

	c, rec := setup1("/offers/new", body)

	if assert.NoError(t, login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Get and set users session token
		body := rec.Body.String()
		var userForm forms.User
		if err := json.Unmarshal([]byte(body), &userForm); err != nil {
			panic(err)
		}
		sessionToken = userForm.Token
	}
}

func TestLoginWithWrongPassword(t *testing.T) {

	body := `{"email":"` + userEmail + `", "password":"abc"}`

	c, rec := setup1("/offers/new", body)

	if assert.NoError(t, login(c)) {
		assert.NotEqual(t, http.StatusOK, rec.Code)
	}
}

func TestGetOffers(t *testing.T) {

	c, rec := setup2("/offers/new", "")

	if assert.NoError(t, getOffers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestNewOffer(t *testing.T) {

	// Should succed
	c, rec := setup2("/offers/new", `{"name":"Flashy Car"}`)
	if assert.NoError(t, newOffer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	// Name is too short - should fail
	c, rec = setup2("/offers/new", `{"name":"Ca"}`)
	if assert.NoError(t, newOffer(c)) {
		assert.NotEqual(t, http.StatusOK, rec.Code)
	}

	// Name is too long - should fail
	c, rec = setup2("/offers/new", `{"name":"Carssssssssssssssssssssssssssssssssssssssssss"}`)
	if assert.NoError(t, newOffer(c)) {
		assert.NotEqual(t, http.StatusOK, rec.Code)
	}
}

// This will test logout, and also logout the user
func TestLogout(t *testing.T) {

	c, rec := setup2("/logout", "")

	if assert.NoError(t, logout(c)) {
		assert.NotEqual(t, http.StatusOK, rec.Code)
	}
}

// When not logged in
func setup1(path string, body string) (echo.Context, *httptest.ResponseRecorder) {

	bodyReader := strings.NewReader(body)

	e := echo.New()

	request := httptest.NewRequest(http.MethodPost, "/", bodyReader)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	c.SetPath(path)

	return c, recorder
}

// When logged in
func setup2(path string, body string) (echo.Context, *httptest.ResponseRecorder) {

	bodyReader := strings.NewReader(body)

	e := echo.New()

	request := httptest.NewRequest(http.MethodPost, "/", bodyReader)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	request.Header.Set(echo.HeaderAuthorization, "Bearer "+sessionToken+" "+userEmail)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	c.SetPath(path)

	return c, recorder
}

// Setup the test database IF it hasn't already been setup with a previous call
func useTestDatabase() {

	if database.Client == nil {
		var err error
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		database.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
		if err != nil {
			panic(err)
		}
		database.MAIN_DATABASE = database.TEST_DATABASE
		database.USER_DATABASE = database.TEST_DATABASE
	}

}
