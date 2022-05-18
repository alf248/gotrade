package database

import (
	"context"
	"errors"
	"net/http"

	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActionForm struct {
	Action string `json:"action" bson:"action"`
}

const (
	Start  = "start"
	Stop   = "stop"
	Delete = "delete"
	Buy    = "buy"
)

const (
	ACTIVE   = "active"
	FINISHED = "finished"
)

func OfferAction(objectId primitive.ObjectID, user *forms.User, aform *ActionForm) (int, error) {

	coll := Client.Database(MAIN_DATABASE).Collection(OFFERS_COLLECTION)

	filter := bson.D{{Key: "_id", Value: objectId}}
	//projection := bson.D{{"type", 1}, {"rating", 1}, {"_id", 0}}
	//opts := options.FindOne().SetProjection(projection)
	var vector forms.Offer
	err := coll.FindOne(context.TODO(), filter).Decode(&vector)
	if err != nil {
		return http.StatusNotFound, err
	}

	switch aform.Action {
	case Buy:
		if vector.Status == ACTIVE || vector.Status == FINISHED {
			return http.StatusForbidden, errors.New("Vector can't be active or finished")
		}
		if vector.Giver == user.Name {
			return http.StatusForbidden, errors.New("Can't buy from self")
		}
		// update receiver field on vector
		// change vector status to "active"
		filter := bson.D{{Key: "_id", Value: objectId}}
		update := bson.D{{"$set", bson.D{{RECEIVER, user.Name}, {"status", ACTIVE}}}}
		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return http.StatusNotFound, err
		}
		if result.ModifiedCount < 0 {
			return http.StatusNotFound, errors.New("Failed to modify")
		}

	case Delete:
		if vector.Status == ACTIVE || vector.Status == FINISHED {
			return http.StatusForbidden, errors.New("Vector can't be active or finished")
		}
		if !(vector.Giver == user.Name || vector.Receiver == user.Name) {
			return http.StatusForbidden, errors.New("Can only delete your own vectors")
		}
		filter := bson.D{{Key: "_id", Value: objectId}}
		result, err := coll.DeleteOne(context.TODO(), filter)
		if err != nil {
			return http.StatusNotFound, err
		}
		if result.DeletedCount < 1 {
			return http.StatusNotFound, errors.New("Could not delete")
		}

	default:
	}

	return http.StatusOK, nil
}
