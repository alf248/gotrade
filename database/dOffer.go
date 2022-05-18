package database

import (
	"context"
	"errors"
	"net/http"

	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewOffer(user *forms.User, form *forms.NewOfferForm) (string, error) {

	coll := Client.Database(MAIN_DATABASE).Collection(OFFERS_COLLECTION)

	result, err := coll.InsertOne(context.TODO(), form)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetOffer(id string) (*forms.Offer, int, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	coll := Client.Database(MAIN_DATABASE).Collection(OFFERS_COLLECTION)

	filter := bson.D{{Key: "_id", Value: objectId}}
	var offer forms.Offer
	err = coll.FindOne(context.TODO(), filter).Decode(&offer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("nothing found")
		}
		return nil, http.StatusInternalServerError, err
	}

	return &offer, 0, nil
}

func GetOffersById(vOfferIds *[]string) ([]forms.Offer, int, error) {

	coll := Client.Database(MAIN_DATABASE).Collection(OFFERS_COLLECTION)

	var ordersObjectID []primitive.ObjectID
	for _, t := range *vOfferIds {
		objectId, err := primitive.ObjectIDFromHex(t)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
		ordersObjectID = append(ordersObjectID, objectId)
	}

	filter := bson.D{{"_id", bson.M{"$in": ordersObjectID}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	var results []forms.Offer
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return results, http.StatusOK, nil
}

func EditOffer(id primitive.ObjectID, user *forms.User, newForm *forms.NewOfferForm) (httpStatus int, e error) {

	coll := Client.Database(MAIN_DATABASE).Collection(OFFERS_COLLECTION)

	// GET THE VECTOR FIRST SO IT CAN BE INSPECTED
	var offer forms.Offer
	filter := bson.D{{Key: "_id", Value: id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&offer)
	if err != nil {
		return http.StatusNotFound, errors.New("")
	}

	// Make sure the user is giver or receiver
	if !(offer.Giver == user.Name || offer.Receiver == user.Name) {
		return http.StatusForbidden, errors.New("user is not in offer")
	}

	// Make sure the offer is not active
	if offer.Status == "active" {
		return http.StatusForbidden, errors.New("can not edit an active offer")
	}

	// IF OK, MAKE THE UPDATE
	bytes, err := bson.Marshal(newForm)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	var updateVal bson.M
	err = bson.Unmarshal(bytes, &updateVal)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	update := bson.D{{Key: "$set", Value: updateVal}}
	res, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return http.StatusNotFound, err
	}

	if res.ModifiedCount < 1 {
		return http.StatusBadRequest, errors.New("nothing modified")
	}

	return http.StatusOK, nil
}
