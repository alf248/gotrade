package forms

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `json:"-" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Token       string             `json:"token" bson:"token"`
	Phone       string             `json:"phone" bson:"phone"`
	Password    string             `json:"-" bson:"password"`
	LoginTime   string             `json:"loginTime" bson:"loginTime"`
	OffersMade  int                `json:"offersMade" bson:"offersMade"`
	VectorsMade int                `json:"vectorsMade" bson:"vectorsMade"`
}

type UserInOut struct {
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	Token       string `json:"token" bson:"token"`
	Phone       string `json:"phone" bson:"phone"`
	Password    string `json:"-" bson:"password"`
	VectorsMade int    `json:"vectorsMade" bson:"vectorsMade"`
	OffersMade  int    `json:"offersMade" bson:"offersMade"`
}

func (u *UserInOut) MakePrivate() {
	u.Email = "***"
	u.Phone = "***"
}

type EditUser struct {
	Phone       string `json:"phone,omitempty" bson:"phone,omitempty"`
	Email       string `json:"-" bson:"email,omitempty"`
	VectorsMade int    `json:"-" bson:"vectorsMade,omitempty"`
	OffersMade  int    `json:"-" bson:"offersMade,omitempty"`
}
