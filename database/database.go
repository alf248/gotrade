package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

var UseAtlasSearch = false

var MAIN_DATABASE = "gotrade"
var USER_DATABASE = MAIN_DATABASE

const TEST_DATABASE = "gotradetest"

const OFFERS_COLLECTION = "offers"
const ORDERS_COLLECTION = "orders"
const USERS_COLLECTION = "users"

const MAX_OFFERS_PER_USER = 100
const MAX_VECTORS_PER_USER = 100

const ORDER_GIVER_FID = "giverFID"
const ORDER_RECEIVER_FID = "receiverFID"

const OFFER_CREATOR_FID = "creatorFID"
const OFFER_ACCEPTOR_FID = "acceptorFID"

const SEARCH_OFFERS_MAX_PAGE_SIZE = 10

func DropOffersCollection() {

	database := Client.Database("test")
	offersCollection := database.Collection("offers")
	//usersCollection := database.Collection("users")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err := offersCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}
