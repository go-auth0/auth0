package management_test

import (
	"fmt"
	"os"

	"github.com/yieldr/go-auth0/management"
)

var (
	domain = os.Getenv("AUTH0_DOMAIN")
	id     = os.Getenv("AUTH0_CLIENT_ID")
	secret = os.Getenv("AUTH0_CLIENT_SECRET")
)

func ExampleUser() {
	m, err := management.New(domain, id, secret)
	if err != nil {
		fmt.Printf("Failed creating management client. %s", err)
	}

	u := &management.User{
		Connection: "Username-Password-Authentication",
		Email:      "smith@example.com",
		Password:   "F4e3DA1a6cDD",
	}

	err = m.User.Create(u)
	if err != nil {
		fmt.Printf("Failed creating user. %s", err)
	}
	defer m.User.Delete(u.ID)

	fmt.Printf("User created!")
	// Output: User created!
}
