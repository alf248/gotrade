package forms

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `json:"-" bson:"_id"`
	FID        string             `json:"fid" bson:"fid"`
	Name       string             `json:"name" bson:"name"`
	OffersMade int                `json:"offersMade" bson:"offersMade"`
	OrdersMade int                `json:"ordersMade" bson:"ordersMade"`
}

type NewUser struct {
	FID  string `bson:"fid"`
	Name string `bson:"name"`
}

type EditUser struct {
	OffersMade  int `bson:"offersMade,omitempty"`
	VectorsMade int `bson:"ordersMade,omitempty"`
}
