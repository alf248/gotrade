package mock

import (
	"context"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"
)

func addMockUsers() error {

	coll := database.Client.Database(database.USER_DATABASE).Collection(database.USERS_COLLECTION)

	var users []any

	for _, v := range UserNames {
		users = append(users, createMockUser(v))
	}

	_, err := coll.InsertMany(context.TODO(), users)
	if err != nil {
		return err
	}

	return nil
}

func createMockUser(name string) *forms.UserInOut {

	user := forms.UserInOut{Name: name, Email: name + "@example.com", Password: "pass"}
	hash, err := database.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hash
	return &user
}
