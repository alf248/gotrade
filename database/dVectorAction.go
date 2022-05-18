package database

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteVector(orderId string, username string) (httpStatus int, e error) {

	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return http.StatusBadRequest, err
	}

	coll := Client.Database(MAIN_DATABASE).Collection(VECTORS_COLLECTION)

	_ = objectId
	filter := bson.D{{"_id", objectId}, {"head", username}}

	r, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return http.StatusNotFound, err
	}
	if r.DeletedCount < 1 {
		return http.StatusNotFound, errors.New("failed to delete")
	}

	return http.StatusOK, nil
}
