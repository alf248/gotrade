package mock

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"
)

func addSomeMockOffers() error {

	coll := database.Client.Database(database.MAIN_DATABASE).Collection(database.OFFERS_COLLECTION)

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	form1 := createRandomOffer(rand)
	form2 := createRandomOffer(rand)

	form1.Name = "Lucky Car"
	form1.Giver = "joe"
	form1.Price = 1000
	form1.Cat = "car"
	form1.Description = "This car brings luck to the rider. Good for poker players."
	form1.Image = "car1"

	form2.Name = "Fast Car"
	form2.Giver = "mia"
	form2.Price = 10000
	form2.Cat = "car"
	form2.Description = "This car goes twice as fast as the speed limit. Bring some cash to pay the cops."
	form2.Image = "car2"

	var forms = []any{form1, form2}

	_, err := coll.InsertMany(context.TODO(), forms)
	if err != nil {
		return err
	}

	return nil
}

func addMockOffers(count int) error {

	coll := database.Client.Database(database.MAIN_DATABASE).Collection(database.OFFERS_COLLECTION)

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	var forms = []any{} // must use interface because mongo wont accept anything else here

	for i := 0; i < count; i++ {
		form := createRandomOffer(rand)
		forms = append(forms, form)
	}

	_, err := coll.InsertMany(context.TODO(), forms)
	if err != nil {
		return err
	}

	return nil
}

func createRandomOffer(r *rand.Rand) forms.Offer {

	var pre = title1[r.Intn(len(title1))]
	var cat = title2[r.Intn(len(title2))]
	var title = pre + " " + cat

	var p = r.Float64() * 10000
	var priceStr = strconv.FormatFloat(p, 'f', 2, 64)
	var price, _ = strconv.ParseFloat(priceStr, 64)

	form := forms.Offer{
		Name:        title,
		Description: loremIpsumLong,
		Giver:       UserNames[r.Intn(len(UserNames))],
		Price:       price,
		Currency:    "Euro",
		Sale:        true,
		Image:       strings.ToLower(cat) + fmt.Sprint(1+rand.Intn(2)),
		Likes:       r.Intn(100),
		Dislikes:    r.Intn(10),
		Cat:         strings.ToLower(cat),
	}

	return form
}
