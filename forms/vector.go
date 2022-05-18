package forms

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ACTIVE = "active"

type VectorInOut struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Giver       string             `json:"tail" bson:"tail,omitempty"`
	Receiver    string             `json:"head" bson:"head,omitempty"`
	Path        string             `json:"path" bson:"path,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Price       float64            `json:"price" bson:"price,omitempty"`
	Currency    string             `json:"currency" bson:"currency,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
	Visibility  string             `json:"visibility" bson:"visibility,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Creator     string             `json:"creator" bson:"creator,omitempty"`
	Created     string             `json:"created" bson:"created,omitempty"`
	Delivery    string             `json:"delivery" bson:"delivery,omitempty"`
	Payment     string             `json:"payment" bson:"payment,omitempty"`
	OfferId     primitive.ObjectID `json:"gunId" bson:"gunId,omitempty"`
	//Parent      string             `json:"parent" bson:"parent,omitempty"`
}

func (v *VectorInOut) Init(o Offer, username string) {
	v.Name = o.Name
	v.Giver = o.Giver
	v.Receiver = o.Receiver
	v.Path = o.Path
	v.Description = o.Description
	v.Price = o.Price
	v.Currency = o.Currency
	v.Image = o.Image
	v.Visibility = o.Visibility
	v.Status = ACTIVE // set active on init
	v.Delivery = "ordered"
	v.Created = time.Now().Format(time.RFC3339)
	v.Creator = username
	v.OfferId = o.Id
}

func (v *VectorInOut) AddUserAsHeadOrTail(username string) error {

	if v.Giver == "" {
		v.Giver = username
		return nil
	} else if v.Receiver == "" {
		v.Receiver = username
		return nil
	}

	return errors.New("either tail or head must be empty")
}
