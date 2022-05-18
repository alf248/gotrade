package database

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func SearchOffers(form *forms.SearchOffers) ([]forms.Offer, int, error) {

	coll := Client.Database(MAIN_DATABASE).Collection(OFFERS_COLLECTION)

	sort := 1
	if !form.SortUp {
		sort = -1
	}

	var pipeline mongo.Pipeline

	if form.Search != "" && UseAtlasSearch {
		searchStage := getSearchStage(form)
		pipeline = append(pipeline, searchStage)
	}

	matchStage := getMatchStage(form)

	sortStage := bson.D{{"$sort", bson.D{{"price", sort}}}}
	skipStage := bson.D{{"$skip", int64(form.Page * SEARCH_OFFERS_MAX_PAGE_SIZE)}}
	limitStage := bson.D{{"$limit", SEARCH_OFFERS_MAX_PAGE_SIZE}}

	pipeline = append(pipeline, matchStage, sortStage, skipStage, limitStage)

	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Println("database error:", err.Error())
		return nil, http.StatusInternalServerError, errors.New("database error: aggregation")
	}

	var offers []forms.Offer
	if err = cursor.All(context.TODO(), &offers); err != nil {
		log.Println("database error:", err.Error())
		return nil, http.StatusInternalServerError, errors.New("database error: conversion")
	}

	return offers, http.StatusOK, nil
}

func getSearchStage(form *forms.SearchOffers) primitive.D {
	// https://www.mongodb.com/docs/atlas/atlas-search/query-syntax/#mongodb-pipeline-pipe.-search
	// https://www.mongodb.com/docs/atlas/atlas-search/operators-and-collectors/#std-label-operators-ref
	// https://www.mongodb.com/docs/atlas/atlas-search/text/#std-label-text-ref

	searchStage := bson.D{
		{Key: "$search", Value: bson.D{
			//{"index", "<index name>"}, // use this if the index does not have a 'default' name
			{Key: "text", Value: bson.D{
				{"query", form.Search},
				//{"path", "name"}, // todo: should be an array to include other fields
				{"path", bson.A{"name", "description"}},
			}},
		}},
	}
	return searchStage
}

func getMatchStage(form *forms.SearchOffers) primitive.D {

	var matchStage bson.D
	var match bson.D

	if form.Sale {
		if form.By == "" {
			match = bson.D{{GIVER, bson.D{bson.E{"$exists", true}, bson.E{"$ne", ""}}}}
		} else {
			match = bson.D{{GIVER, form.By}}
		}
	} else {
		if form.By == "" {
			match = bson.D{{RECEIVER, bson.D{bson.E{"$exists", true}, bson.E{"$ne", ""}}}}
		} else {
			match = bson.D{{RECEIVER, form.By}}
		}
	}

	if form.Cat != "" {
		match = append(match, bson.E{"cat", form.Cat})
	}

	matchStage = bson.D{{"$match", match}}

	return matchStage
}
