package management

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/dapperlabs/auth0/v4"
	"github.com/dapperlabs/auth0/v4/internal/testing/expect"
)

func TestUser(t *testing.T) {

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
		Password:   auth0.String("Passwords hide their Chuck"),
		Username:   auth0.String("chucknorris"),
		GivenName:  auth0.String("Chuck"),
		FamilyName: auth0.String("Norris"),
		Nickname:   auth0.String("Chucky"),
		UserMetadata: map[string]interface{}{
			"favourite_attack": "roundhouse_kick",
		},
		EmailVerified: ConvertibleBool(true),
		VerifyEmail:   auth0.Bool(false),
		AppMetadata: map[string]interface{}{
			"facts": []string{
				"count_to_infinity_twice",
				"kill_two_stones_with_one_bird",
				"can_hear_sign_language",
				"knows_victorias_secret",
			},
		},
		Picture: auth0.String("https://example-picture-url.jpg"),
		Blocked: auth0.Bool(false),
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

	s := &ResourceServer{
		Name: auth0.Stringf("Test Role (%s)",
			time.Now().Format(time.StampMilli)),
		Identifier: auth0.String("https://api.example.com/role"),
		Scopes: []*ResourceServerScope{
			{
				Value:       auth0.String("read:resource"),
				Description: auth0.String("Read Resource"),
			},
			{
				Value:       auth0.String("update:resource"),
				Description: auth0.String("Update Resource"),
			},
		},
	}
	err = m.ResourceServer.Create(s)
	if err != nil {
		t.Fatal(err)
	}
	defer m.ResourceServer.Delete(auth0.StringValue(s.ID))

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
		ul, err := m.User.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", ul)
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

		t.Run("AppMetadata", func(t *testing.T) {
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
	})

	t.Run("Roles", func(t *testing.T) {
		l, err := m.User.Roles(auth0.StringValue(u.ID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l.Roles)
	})

	t.Run("AssignRoles", func(t *testing.T) {
		err = m.User.AssignRoles(auth0.StringValue(u.ID), r1, r2)
		if err != nil {
			t.Error(err)
		}

	})

	t.Run("RemoveRoles", func(t *testing.T) {
		roles := []*Role{r1, r2}
		err = m.User.RemoveRoles(auth0.StringValue(u.ID), roles...)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Permissions", func(t *testing.T) {
		l, err := m.User.Permissions(auth0.StringValue(u.ID))
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", l.Permissions)
	})

	t.Run("AssignPermissions", func(t *testing.T) {
		permissions := []*Permission{
			{Name: auth0.String("read:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
			{Name: auth0.String("update:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
		}
		err = m.User.AssignPermissions(auth0.StringValue(u.ID), permissions...)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("RemovePermissions", func(t *testing.T) {
		permissions := []*Permission{
			{Name: auth0.String("read:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
			{Name: auth0.String("update:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
		}
		err = m.User.RemovePermissions(auth0.StringValue(u.ID), permissions...)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Blocks", func(t *testing.T) {
		b, err := m.User.Blocks(u.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", b)
	})

	t.Run("Blocks", func(t *testing.T) {
		err := m.User.Unblock(u.GetID())
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.User.Delete(u.GetID())
		if err != nil {
			t.Fatal(err)
		}
	})

	// Create some users we can search for
	allUsers := []*User{
		{
			Email:      auth0.String("alice@example.com"),
			Username:   auth0.String("alice"),
			Password:   auth0.String("5301111b-b31b-47c4-bf3d-0c26ea57bdf4"),
			Connection: auth0.String("Username-Password-Authentication"),
		},
		{
			Email:      auth0.String("bob@example.com"),
			Username:   auth0.String("bob"),
			Password:   auth0.String("bcfc3bca-8cd3-4b74-a474-402420f34f85"),
			Connection: auth0.String("Username-Password-Authentication"),
		},
		{
			Email:      auth0.String("charlie@example.com"),
			Username:   auth0.String("charlie"),
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
	defer func() {
		for _, user := range allUsers {
			m.User.Delete(auth0.StringValue(user.ID))
		}
	}()

	t.Run("Search", func(t *testing.T) {
		ul, err := m.User.Search(Query(`email:"alice@example.com"`))
		if err != nil {
			t.Error(err)
		}
		if len(ul.Users) != 1 {
			t.Error("unexpected number of users found")
		}
		t.Logf("%v\n", ul)
	})

	t.Run("ListByEmail", func(t *testing.T) {
		us, err := m.User.ListByEmail("alice@example.com")
		if err != nil {
			t.Error(err)
		}
		if len(us) != 1 {
			t.Error("unexpected number of users found")
		}
		t.Logf("%v\n", us)
	})
}

func TestUserIdentity(t *testing.T) {

	t.Run("MarshalJSON", func(t *testing.T) {
		for u, expected := range map[*UserIdentity]string{
			&UserIdentity{}:                            `{}`,
			&UserIdentity{UserID: auth0.String("1")}:   `{"user_id":"1"}`,
			&UserIdentity{UserID: auth0.String("foo")}: `{"user_id":"foo"}`,
			&UserIdentity{UserID: auth0.String("foo"),
				ProfileData: ProfileData{
					Email: auth0.String("test@example.com"),
				}}: `{"user_id":"foo", "profile_data":{"email":"test@example.com"}}`,
		} {
			b, err := json.Marshal(u)
			if err != nil {
				t.Error(err)
			}
			expect.Expect(t, string(b), expected)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		for b, expected := range map[string]*UserIdentity{
			`{}`:                &UserIdentity{UserID: nil},
			`{"user_id":1}`:     &UserIdentity{UserID: auth0.String("1")},
			`{"user_id":"1"}`:   &UserIdentity{UserID: auth0.String("1")},
			`{"user_id":"foo"}`: &UserIdentity{UserID: auth0.String("foo")},
			`{"user_id":"foo", "profile_data":{"email":"test@example.com"}}`: &UserIdentity{UserID: auth0.String("foo"),
				ProfileData: ProfileData{
					Email: auth0.String("test@example.com"),
				}},
		} {
			var u UserIdentity
			err := json.Unmarshal([]byte(b), &u)
			if err != nil {
				t.Error(err)
			}
			expect.Expect(t, u.GetUserID(), expected.GetUserID())
		}
	})
}

func TestConvertibleBoolean(t *testing.T) {

	t.Run("MarshalJSON", func(t *testing.T) {
		for _, test := range []struct {
			in       *ConvertibleBoolean
			expected []byte
		}{
			{ConvertibleBool(false), []byte(`{"bool":false}`)},
			{ConvertibleBool(true), []byte(`{"bool":true}`)},
		} {
			var ts struct {
				Bool *ConvertibleBoolean `json:"bool,omitempty"`
			}
			ts.Bool = test.in
			str, err := json.Marshal(&ts)
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}

			if !bytes.Equal(str, test.expected) {
				t.Errorf("unexpected output. have %v, expected %v", ts.Bool, test.expected)
			}
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		for _, test := range []struct {
			in       string
			expected *ConvertibleBoolean
		}{
			{`{"bool": false}`, ConvertibleBool(false)},
			{`{"bool": "false"}`, ConvertibleBool(false)},
			{`{"bool": true}`, ConvertibleBool(true)},
			{`{"bool": "true"}`, ConvertibleBool(true)},
		} {
			var ts struct {
				Bool *ConvertibleBoolean `json:"bool,omitempty"`
			}
			err := json.Unmarshal([]byte(test.in), &ts)
			if err != nil {
				t.Errorf("expected no error, got %s", err)
			}

			if ConvertibleBoolValue(ts.Bool) != ConvertibleBoolValue(test.expected) {
				t.Errorf("unexpected output. have %v, expected %v", ts.Bool, test.expected)
			}
		}
	})

	t.Run("ConvertibleBoolValue", func(t *testing.T) {
		for _, test := range []struct {
			in       *ConvertibleBoolean
			expected ConvertibleBoolean
		}{
			{nil, false},
			{ConvertibleBool(false), false},
			{ConvertibleBool(true), true},
		} {
			have := ConvertibleBoolValue(test.in)
			if have != test.expected {
				t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
			}
		}

	})
}
