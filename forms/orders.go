package forms

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ACTIVE = "active"

// A vector is synonymous with "order"
type Order struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`

	GiverFID  string             `json:"giverFID,omitempty" bson:"giverFID,omitempty"`
	GiverID   primitive.ObjectID `json:"giverID,omitempty" bson:"giverID,omitempty"`
	GiverName string             `json:"giverName,omitempty" bson:"giverName,omitempty"`

	ReceiverFID  string             `json:"receiverFID,omitempty" bson:"receiverFID,omitempty"`
	ReceiverID   primitive.ObjectID `json:"receiverID,omitempty" bson:"receiverID,omitempty"`
	ReceiverName string             `json:"receiverName,omitempty" bson:"receiverName,omitempty"`

	Path        string `json:"path,omitempty" bson:"path,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Category    string `json:"category,omitempty" bson:"category,omitempty"`

	Price    float64 `json:"price" bson:"price,omitempty"`
	Currency string  `json:"currency" bson:"currency,omitempty"`

	Image string `json:"image,omitempty" bson:"image,omitempty"`

	Visibility string `json:"visibility" bson:"visibility,omitempty"`

	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	Created  string             `json:"created,omitempty" bson:"created,omitempty"`
	Delivery string             `json:"delivery,omitempty" bson:"delivery,omitempty"`
	Payment  string             `json:"payment,omitempty" bson:"payment,omitempty"`
	OfferId  primitive.ObjectID `json:"offerId" bson:"offerId,omitempty"`
	//Parent      string             `json:"parent" bson:"parent,omitempty"`
}

func (v *Order) Init(o Offer, acceptor *User) error {

	if o.CreatorFID == acceptor.FID {
		return errors.New("can not accept own offer")
	}

	if o.Sale {
		v.GiverFID = o.CreatorFID
		v.GiverID = o.CreatorID
		v.GiverName = o.CreatorName

		v.ReceiverFID = acceptor.FID
		v.ReceiverID = acceptor.ID
		v.ReceiverName = acceptor.Name
	} else {
		v.GiverFID = acceptor.FID
		v.GiverID = acceptor.ID
		v.GiverName = acceptor.Name

		v.ReceiverFID = o.CreatorFID
		v.ReceiverID = o.CreatorID
		v.ReceiverName = o.CreatorName
	}

	v.Name = o.Name
	v.Path = o.Path
	v.Description = o.Description
	v.Category = o.Category
	v.Price = o.Price
	v.Currency = o.Currency
	v.Image = o.Image
	v.Visibility = o.Visibility
	v.Status = ACTIVE // set active on init
	v.Delivery = "ordered"
	v.Created = time.Now().Format(time.RFC3339)
	v.OfferId = o.ID

	return nil
}

type SearchOrders struct {
	Search  string `json:"search,omitempty"` // a search query
	Max     int    `json:"max,omitempty"`    // the max amount of results
	Page    int    `json:"page,omitempty"`   // search pagination
	SortBy  string `json:"sortBy,omitempty"`
	SortUp  bool   `json:"sortUp,omitempty"`
	Active  bool   `json:"active,omitempty"`
	AsGiver bool   `json:"asGiver"`
}

func (o *SearchOrders) Curate() error {
	return nil
}

type OrderAction struct {
	Action string `json:"action"`
}

func (o *OrderAction) Curate() error {

	actions := []string{"delete", "pause"}

	for _, action := range actions {
		if o.Action == action {
			return nil
		}
	}

	return errors.New("action must be one of " + strings.Join(actions, " "))
}
