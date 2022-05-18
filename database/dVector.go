package database

import (
	"context"
	"net/http"

	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewVectors(vectors []forms.VectorInOut) (httpStatus int, e error) {

	coll := Client.Database(MAIN_DATABASE).Collection(VECTORS_COLLECTION)

	var vectorsAny []any
	for _, v := range vectors {
		vectorsAny = append(vectorsAny, v)
	}

	_, err := coll.InsertMany(context.TODO(), vectorsAny)
	if err != nil {
		return http.StatusConflict, err
	}

	return http.StatusOK, nil
}

func GetVectors(sortUp bool, active bool) ([]forms.VectorInOut, int, error) {

	coll := Client.Database(MAIN_DATABASE).Collection(VECTORS_COLLECTION)

	sortVal := -1
	if sortUp {
		sortVal = 1
	}
	sort := bson.D{{"price", sortVal}}

	filterVal := "active"
	if !active {
		filterVal = "finished"
	}
	filter := bson.D{{"status", filterVal}}

	//projection := bson.D{{"name", 1}, {"price", 1}, {"_id", bson.E{"$toString", "$_id"}}}
	// SetProjection(projection)
	opts := options.Find().SetLimit(10).SetSort(sort)
	// SetSkip(1) .. skips the first document

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	var vectors []forms.VectorInOut
	if err = cursor.All(context.TODO(), &vectors); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return vectors, http.StatusOK, nil
}
