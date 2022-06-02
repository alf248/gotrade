package database

import (
	"context"
	"net/http"

	"github.com/alf248/gotrade/forms"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewOrders(orders []forms.Order) (httpStatus int, e error) {

	coll := Client.Database(MAIN_DATABASE).Collection(ORDERS_COLLECTION)

	var ordersAny []any
	for _, v := range orders {
		ordersAny = append(ordersAny, v)
	}

	_, err := coll.InsertMany(context.TODO(), ordersAny)
	if err != nil {
		return http.StatusConflict, err
	}

	return http.StatusOK, nil
}

func GetOrders(sortUp bool, active bool, asGiver bool, byFID string) ([]forms.Order, int, error) {

	coll := Client.Database(MAIN_DATABASE).Collection(ORDERS_COLLECTION)

	sortVal := -1
	if sortUp {
		sortVal = 1
	}
	sort := bson.D{{"price", sortVal}}

	status := "active"
	if !active {
		status = "finished"
	}
	filter := bson.D{{"status", status}}

	if asGiver {
		filter = append(filter, bson.E{ORDER_GIVER_FID, byFID})
	} else {
		filter = append(filter, bson.E{ORDER_RECEIVER_FID, byFID})
	}

	opts := options.Find().SetLimit(20).SetSort(sort)

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	var orders []forms.Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	//fmt.Printf("orders: %+v", orders)

	return orders, http.StatusOK, nil
}
