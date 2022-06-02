package forms

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Offer struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Name        string `json:"name" bson:"name,omitempty"`
	Path        string `json:"path" bson:"path,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Category    string `json:"cat,omitempty" bson:"cat,omitempty"`

	Sale bool `json:"sale" bson:"sale"`

	CreatorFID   string             `json:"creatorFID" bson:"creatorFID,omitempty"`
	CreatorID    primitive.ObjectID `json:"creatorID" bson:"creatorID,omitempty"`
	CreatorName  string             `json:"creatorName" bson:"creatorName,omitempty"`
	AcceptorFID  string             `json:"acceptorFID" bson:"acceptorFID,omitempty"`
	AcceptorID   primitive.ObjectID `json:"acceptorID" bson:"acceptorID,omitempty"`
	AcceptorName string             `json:"acceptorName" bson:"acceptorName,omitempty"`

	Price    float64 `json:"price" bson:"price,omitempty"`
	Currency string  `json:"currency" bson:"currency,omitempty"`

	Image string `json:"image" bson:"image,omitempty"`

	Visibility string `json:"visibility" bson:"visibility,omitempty"`
	Status     string `json:"status" bson:"status,omitempty"`
	Charges    int    `json:"charges" bson:"charges,omitempty"`

	Likes    int `json:"likes" bson:"likes,omitempty"`
	Dislikes int `json:"dislikes" bson:"dislikes"`
}

type NewOfferForm struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description,omitempty"`
	Category    string  `json:"category" bson:"category"`
	Sale        bool    `json:"sale" bson:"sale,omitempty"`
	Price       float64 `json:"price" bson:"price"`
	Currency    string  `json:"currency" bson:"currency"`
	Image       string  `json:"image" bson:"image,omitempty"`
}

func (f *NewOfferForm) CreateOffer(user *User) (Offer, error) {

	o := Offer{
		Name:        f.Name,
		Description: f.Description,
		Sale:        f.Sale,
		Price:       f.Price,
		Currency:    f.Currency,
		Image:       f.Image,

		CreatorFID:  user.FID,
		CreatorID:   user.ID,
		CreatorName: user.Name,
	}

	return o, nil
}

// Verify and modify the form as needed
func (f *NewOfferForm) Curate() error {

	if utf8.RuneCountInString(f.Name) > 20 {
		return errors.New("name is too long; over 20, got" + fmt.Sprintf("%d", utf8.RuneCountInString(f.Name)))
	}

	if utf8.RuneCountInString(f.Name) < 3 {
		return errors.New("name is too short; less than 3")
	}

	if utf8.RuneCountInString(f.Description) > 300 {
		return errors.New("description is too long; over 300")
	}

	if f.Price > 9999999999999999999999999999999999999 {
		return errors.New("price is too big; over 9999999999999999999999999999999999999")
	}

	if f.Price < 0 {
		return errors.New("price is too small; less than 0")
	}

	if utf8.RuneCountInString(f.Currency) >= 10 {
		return errors.New("currency string is too long; over 10")
	}

	if utf8.RuneCountInString(f.Image) >= 100 {
		return errors.New("image string is too long; over 100")
	}

	return nil
}

type SearchOffers struct {
	Search   string `json:"search,omitempty"`
	Sale     bool   `json:"sale,omitempty"`
	ByFID    string `json:"by,omitempty"`     // by which user
	Max      int64  `json:"max,omitempty"`    // max returned results
	SortBy   string `json:"sortBy,omitempty"` // price
	SortUp   bool   `json:"sortUp,omitempty"`
	Category string `json:"cat,omitempty"` // category
	Page     int    `json:"page,omitempty"`
}

func (form *SearchOffers) Verify() error {
	return nil
}
