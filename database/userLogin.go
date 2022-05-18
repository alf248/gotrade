package database

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(email string, password string) (*forms.User, error) {

	token := GenerateSessionToken(20)

	coll := Client.Database(USER_DATABASE).Collection(USERS_COLLECTION)

	var user forms.User
	err := coll.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("password hash failed")
	}

	now := time.Now().Format(time.RFC3339)

	filter := bson.D{{"email", email}}
	update := bson.D{{"$set", bson.D{{"token", token}, {"loginTime", now}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount < 1 {
		return nil, err
	}

	user.Token = token

	return &user, nil
}

func CheckUser(email string, token string) (*forms.User, int, error) {

	coll := Client.Database(USER_DATABASE).Collection(USERS_COLLECTION)

	var user forms.User

	err := coll.FindOne(context.TODO(), bson.D{{"email", email}, {"token", token}}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, http.StatusUnauthorized, errors.New("")
	}
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("database error")
	}

	// Check that session token has not expired
	then, err := time.Parse(time.RFC3339, user.LoginTime)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("timer error")
	}
	now := time.Now()

	timeDifHours := now.Sub(then).Hours()
	if timeDifHours > 1 {
		return nil, http.StatusUnauthorized, errors.New("login expired")
	}

	return &user, http.StatusOK, nil
}

func Logout(email string, token string) error {
	// remove users token from mongodb
	// signal to user that he has logged out (clear session storage on browser)

	coll := Client.Database(USER_DATABASE).Collection(USERS_COLLECTION)

	filter := bson.D{{"email", email}, {"token", token}}
	update := bson.D{{"$set", bson.D{{"token", ""}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}
	if result.MatchedCount < 1 {
		return errors.New("could not find user with email" + email)
	}
	if result.ModifiedCount < 1 {
		return errors.New("database error")
	}

	return nil
}
