package management_test

import (
	"fmt"
	"os"

	"gopkg.in/auth0.v1"
	"gopkg.in/auth0.v1/management"
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
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("smith@example.com"),
		Password:   auth0.String("F4e3DA1a6cDD"),
	}

	err = m.User.Create(u)
	if err != nil {
		fmt.Printf("Failed creating user. %s", err)
	}
	defer m.User.Delete(auth0.StringValue(u.ID))

	fmt.Printf("User created!")
	// Output: User created!
}
