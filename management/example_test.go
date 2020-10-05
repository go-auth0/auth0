package management_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

var (
	domain = os.Getenv("AUTH0_DOMAIN")
	id     = os.Getenv("AUTH0_CLIENT_ID")
	secret = os.Getenv("AUTH0_CLIENT_SECRET")

	api *management.Management
)

func init() {
	var err error
	api, err = management.New(domain, id, secret)
	if err != nil {
		panic(err)
	}
}

func ExampleNew() {
	api, err := management.New(domain, id, secret)
	if err != nil {
		// handle err
	}
	_, _ = api.Stat.ActiveUsers(context.Background())
}

func ExampleClientManager_Create() {
	c := &management.Client{
		Name:        auth0.String("Example Client"),
		Description: auth0.String("This client was created from the Auth0 SDK examples"),
	}

	ctx := context.Background()

	err := api.Client.Create(ctx, c)
	if err != nil {
		// handle err
	}
	defer api.Client.Delete(ctx, c.GetClientID())

	_ = c.GetClientID()
	_ = c.GetClientSecret() // Generated values are available after creation
}

func ExampleResourceServer_List() {
	l, err := api.ResourceServer.List(context.Background())
	if err != nil {
		// handle err
	}
	_ = l.ResourceServers
}

func ExampleUserManager_Create() {
	ctx := context.Background()

	u := &management.User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("smith@example.com"),
		Username:   auth0.String("smith"),
		Password:   auth0.String("F4e3DA1a6cDD"),
	}

	err := api.User.Create(ctx, u)
	if err != nil {
		// handle err
	}
	defer api.User.Delete(ctx, u.GetID())

	_ = u.GetID()
	_ = u.GetCreatedAt()
}

func ExampleRoleManager_Create() {
	ctx := context.Background()

	r := &management.Role{
		Name:        auth0.String("admin"),
		Description: auth0.String("Administrator"),
	}

	err := api.Role.Create(ctx, r)
	if err != nil {
		// handle err
	}
	defer api.Role.Delete(ctx, r.GetID())
}

var (
	user  = &management.User{}
	admin = &management.Role{}
)

func ExampleUserManager_AssignRoles() {
	ctx := context.Background()

	err := api.User.AssignRoles(ctx, user.GetID(), admin)
	if err != nil {
		// handle err
	}
	defer api.User.RemoveRoles(ctx, user.GetID(), admin)
}

func ExampleUserManager_List() {
	q := management.Query(`name:"jane smith"`)
	l, err := api.User.List(context.Background(), q)
	if err != nil {
		// handle err
	}
	_ = l.Users // users matching name "jane smith"
}

func ExampleUserManager_List_pagination() {
	var page int
	for {
		l, err := api.User.List(
			context.Background(),
			management.Query(`logins_count:{100 TO *]`),
			management.Page(page))
		if err != nil {
			// handle err
		}
		for _, u := range l.Users {
			u.GetID() // do something with each user
		}
		if !l.HasNext() {
			break
		}
		page++
	}
}

func ExampleConnectionManager_List() {
	l, err := api.Connection.List(context.Background(), management.Parameter("strategy", "auth0"))
	if err != nil {
		// handle err
	}
	for _, c := range l.Connections {

		fmt.Println(c.GetName())

		if o, ok := c.Options.(*management.ConnectionOptions); ok {
			fmt.Printf("\tPassword Policy: %s\n", o.GetPasswordPolicy())
			fmt.Printf("\tMulti-Factor Auth Enabled: %t\n", o.MFA["active"])
		}
	}
	// Output: Username-Password-Authentication
	// 	Password Policy: good
	// 	Multi-Factor Auth Enabled: true
}

func ExampleConnectionManager_Create() {
	c := &management.Connection{
		Name:     auth0.Stringf("Test-Google-OAuth2-%d", time.Now().Unix()),
		Strategy: auth0.String("google-oauth2"),
		Options: &management.ConnectionOptionsGoogleOAuth2{
			ClientID:     auth0.String(""), // replace with your client id
			ClientSecret: auth0.String(""),
			AllowedAudiences: []interface{}{
				"example.com",
				"api.example.com",
			},
			Profile:  auth0.Bool(true),
			Calendar: auth0.Bool(true),
			Youtube:  auth0.Bool(false),
		},
	}

	ctx := context.Background()

	defer api.Connection.Delete(ctx, c.GetID())

	err := api.Connection.Create(ctx, c)
	if err != nil {
		// handle err
	}
}
