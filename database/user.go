package database

import (
	"context"
	"errors"
	"net/http"

	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUser(user forms.UserInOut) (string, error) {

	coll := Client.Database(USER_DATABASE).Collection(USERS_COLLECTION)

	hash, err := HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hash

	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetUser(name string) (*forms.UserInOut, int, error) {

	var user forms.UserInOut

	coll := Client.Database(USER_DATABASE).Collection(USERS_COLLECTION)

	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, http.StatusNotFound, errors.New("Not Found")
	}
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("database error")
	}

	return &user, http.StatusOK, nil
}

func EditUser(name string, editForm *forms.EditUser) (*forms.User, int, error) {

	coll := Client.Database(USER_DATABASE).Collection(USERS_COLLECTION)

	pByte, err := bson.Marshal(editForm)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var set bson.M
	err = bson.Unmarshal(pByte, &set)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	filter := bson.D{{"name", name}}
	update := bson.D{{"$set", set}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	result := coll.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, http.StatusNotFound, result.Err()
	}

	var doc forms.User
	decodeErr := result.Decode(&doc)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, decodeErr
	}

	return &doc, http.StatusOK, nil
}
