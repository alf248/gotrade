package forms

import (
	"errors"
	"unicode/utf8"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Offer struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Name        string  `json:"name" bson:"name,omitempty"`
	Path        string  `json:"path" bson:"path,omitempty"`
	Description string  `json:"description" bson:"description,omitempty"`
	Giver       string  `json:"giver" bson:"giver,omitempty"`
	Receiver    string  `json:"receiver" bson:"receiver,omitempty"`
	Sale        bool    `json:"sale" bson:"sale,omitempty"`
	Price       float64 `json:"price" bson:"price,omitempty"`
	Currency    string  `json:"currency" bson:"currency,omitempty"`

	Cat string `json:"cat" bson:"cat,omitempty"`

	Image      string `json:"image" bson:"image,omitempty"`
	Visibility string `json:"visibility" bson:"visibility,omitempty"`
	Status     string `json:"status" bson:"status,omitempty"`
	Charges    int    `json:"charges" bson:"charges,omitempty"`

	Likes    int `json:"likes" bson:"likes,omitempty"`
	Dislikes int `json:"dislikes" bson:"dislikes"`
}

type SearchOffers struct {
	Search string `json:"search,omitempty"`
	Sale   bool   `json:"sale,omitempty"`
	By     string `json:"by,omitempty"`     // by which user
	Max    int64  `json:"max,omitempty"`    // max returned results
	SortBy string `json:"sortBy,omitempty"` // price
	SortUp bool   `json:"sortUp,omitempty"`
	Cat    string `json:"cat,omitempty"` // category
	Page   int    `json:"page,omitempty"`
}

func (form *SearchOffers) Verify() error {
	return nil
}

type NewOfferForm struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description,omitempty"`
	Sale        bool    `json:"sale" bson:"sale,omitempty"`
	Price       float64 `json:"price" bson:"price"`
	Currency    string  `json:"currency" bson:"currency"`
	Image       string  `json:"image" bson:"image,omitempty"`
}

// Verify and modify the form as needed
func (f *NewOfferForm) Curate(user *User) error {

	if utf8.RuneCountInString(f.Name) > 20 {
		return errors.New("name is too long; over 40")
	}

	if utf8.RuneCountInString(f.Name) < 3 {
		return errors.New("name is too short; less than 3")
	}

	if utf8.RuneCountInString(f.Description) > 300 {
		return errors.New("description is too long; over 300")
	}

	if f.Price >= 9999999999999999999999999999999999999 {
		return errors.New("name is too long; over 9999999999999999999999999999999999999")
	}

	if utf8.RuneCountInString(f.Currency) >= 10 {
		return errors.New("currency is too long; over 10")
	}

	if utf8.RuneCountInString(f.Image) >= 100 {
		return errors.New("image string is too long; over 100")
	}

	return nil
}
