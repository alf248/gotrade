package mock

import "log"

var title1 = []string{"Fast", "Fabulous", "Excellent", "Shiny", "Speedy", "Lucky"}
var title2 = []string{"Car", "Boat", "Plane", "Bike"}

var loremIpsumLong = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
var loremIpsumShort = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

var UserNames = []string{"joe", "dan", "tom", "tim", "bob", "mia", "zoe", "lea"}

func AddMockUsers() {

	log.Println("Adding mock users")

	addMockUsers()
}

func AddMockOffers(offers int) {

	log.Println("Adding mock offers")

	addMockOffers(offers)
	addSomeMockOffers()
}
