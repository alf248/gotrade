package mock

import (
	"log"

	"github.com/alf248/gotrade/forms"
)

var title1 = []string{"Fast", "Nice", "Cool", "Plain", "Speedy", "Lucky"}
var title2 = []string{"Car", "Boat", "Plane", "Bike"}

var loremIpsumLong = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
var loremIpsumShort = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

var mockUsers []*forms.NewUser

func AddMockUsers() {
	log.Println("Adding mock users")
	addMockUsers()
}

func AddMockOffers(offers int) {
	log.Println("Adding mock offers")
	addMockOffers(offers)
	addSomeMockOffers()
}
