package database

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOrder(orderId string, receiverFID string) (httpStatus int, e error) {

	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return http.StatusBadRequest, err
	}

	fmt.Println("DELETE got id", orderId)
	fmt.Printf("DELETE oid %+v", objectId)

	println("receiver: ", receiverFID)

	coll := Client.Database(MAIN_DATABASE).Collection(ORDERS_COLLECTION)

	_ = objectId
	filter := bson.D{{"_id", objectId}, {ORDER_RECEIVER_FID, receiverFID}}

	r, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return http.StatusNotFound, err
	}
	if r.DeletedCount < 1 {
		return http.StatusNotFound, errors.New("failed to delete")
	}

	return http.StatusOK, nil
}
