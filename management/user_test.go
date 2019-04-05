package management

import (
	"testing"

	"gopkg.in/auth0.v1"
)

func TestUser(t *testing.T) {

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
		Password:   auth0.String("Passwords hide their Chuck"),
		Nickname:   auth0.String("Chucky"),
		UserMetadata: map[string]interface{}{
			"favourite_attack": "roundhouse_kick",
		},
		EmailVerified: auth0.Bool(true),
		VerifyEmail:   auth0.Bool(false),
		AppMetadata: map[string]interface{}{
			"facts": []string{
				"count_to_infinity_twice",
				"kill_two_stones_with_one_bird",
				"can_hear_sign_language",
				"knows_victorias_secret",
			},
		},
	}

	var err error

	r1 := &Role{
		Name:        auth0.String("admin"),
		Description: auth0.String("Administrator"),
	}
	err = m.Role.Create(r1)
	if err != nil {
		t.Fatal(err)
	}

	r2 := &Role{
		Name:        auth0.String("user"),
		Description: auth0.String("User"),
	}
	err = m.Role.Create(r2)
	if err != nil {
		t.Fatal(err)
	}

	defer m.Role.Delete(auth0.StringValue(r1.ID))
	defer m.Role.Delete(auth0.StringValue(r2.ID))

	t.Run("Create", func(t *testing.T) {
		err = m.User.Create(u)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", u)
	})

	t.Run("Read", func(t *testing.T) {
		u, err = m.User.Read(auth0.StringValue(u.ID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", u)
	})

	t.Run("List", func(t *testing.T) {
		us, err := m.User.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", us)
	})

	t.Run("Update", func(t *testing.T) {
		uu := &User{
			Connection: auth0.String("Username-Password-Authentication"),
			Password:   auth0.String("I don't need one"),
		}
		err = m.User.Update(auth0.StringValue(u.ID), uu)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", uu)
	})

	t.Run("Update App Metadata", func(t *testing.T) {
		uu := &User{
			Connection: auth0.String("Username-Password-Authentication"),
			AppMetadata: map[string]interface{}{
				"foo": "bar",
			},
		}
		err = m.User.Update(auth0.StringValue(u.ID), uu)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", uu)
	})

	t.Run("GetRoles", func(t *testing.T) {
		var roles []*Role
		roles, err = m.User.GetRoles(auth0.StringValue(u.ID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", roles)
	})

	t.Run("AssignRoles", func(t *testing.T) {
		roles := []*Role{r1, r2}
		err = m.User.AssignRoles(auth0.StringValue(u.ID), roles...)
		if err != nil {
			t.Error(err)
		}

	})

	t.Run("UnassignRoles", func(t *testing.T) {
		roles := []*Role{r1, r2}
		err = m.User.UnassignRoles(auth0.StringValue(u.ID), roles...)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.User.Delete(auth0.StringValue(u.ID))
	})

	t.Run("Search", func(t *testing.T) {

		// Create some users we can search for
		allUsers := []*User{
			{
				Email:      auth0.String("alice@example.com"),
				Password:   auth0.String("5301111b-b31b-47c4-bf3d-0c26ea57bdf4"),
				Connection: auth0.String("Username-Password-Authentication"),
			},
			{
				Email:      auth0.String("bob@example.com"),
				Password:   auth0.String("bcfc3bca-8cd3-4b74-a474-402420f34f85"),
				Connection: auth0.String("Username-Password-Authentication"),
			},
			{
				Email:      auth0.String("charlie@example.com"),
				Password:   auth0.String("80140c2a-b5c1-490c-a4bf-b0623114d5fd"),
				Connection: auth0.String("Username-Password-Authentication"),
			},
		}
		for _, user := range allUsers {
			err = m.User.Create(user)
			if err != nil {
				t.Error(err)
			}
		}

		// Now search for one of those
		foundUsers, err := m.User.Search(
			Parameter("q", `email:"alice@example.com"`),
			Parameter("search_engine", "v3"))

		if err != nil {
			t.Error(err)
		}
		if len(foundUsers) != 1 {
			t.Error("unexpected number of users found")
		}
		t.Logf("%v\n", foundUsers)

		// Finally clean up
		for _, user := range allUsers {
			m.User.Delete(auth0.StringValue(user.ID))
		}
	})
}
