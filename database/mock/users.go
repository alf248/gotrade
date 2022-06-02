package mock

import (
	"context"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/forms"
)

func addMockUsers() error {

	coll := database.Client.Database(database.USER_DATABASE).Collection(database.USERS_COLLECTION)

	var users []any

	joe := forms.NewUser{Name: "joe", FID: "Xh1RXSzRYqVzJ3cZVktb7X1Ec9x1"}
	tester := forms.NewUser{Name: "tester", FID: "X6pLU9PK5XObMKB6i6VimIHJPIn1"}

	mockUsers = append(mockUsers, &joe)
	mockUsers = append(mockUsers, &tester)
	users = append(users, joe)
	users = append(users, tester)

	_, err := coll.InsertMany(context.TODO(), users)
	if err != nil {
		return err
	}

	return nil
}
