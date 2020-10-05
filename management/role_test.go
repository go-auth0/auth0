package management

import (
	"testing"
	"time"

	"gopkg.in/auth0.v5"
)

func TestRole(t *testing.T) {

	var err error

	r := &Role{
		Name:        auth0.String("admin"),
		Description: auth0.String("Administrator"),
	}

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
		Username:   auth0.String("chuck"),
		Password:   auth0.String("Passwords hide their Chuck"),
	}

	err = m.User.Create(mctx, u)
	if err != nil {
		t.Error(err)
	}
	defer m.User.Delete(mctx, u.GetID())

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
	err = m.ResourceServer.Create(mctx, s)
	if err != nil {
		t.Fatal(err)
	}
	defer m.ResourceServer.Delete(mctx, s.GetID())

	t.Run("Create", func(t *testing.T) {
		err = m.Role.Create(mctx, r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.Role.Read(mctx, r.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Update", func(t *testing.T) {
		id := auth0.StringValue(r.ID)

		r.ID = nil // read-only
		r.Description = auth0.String("The Administrator")

		err = m.Role.Update(mctx, id, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("List", func(t *testing.T) {
		var rs *RoleList
		rs, err = m.Role.List(mctx)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", rs)
	})

	t.Run("AssignUsers", func(t *testing.T) {
		err = m.Role.AssignUsers(mctx, r.GetID(), u)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Users", func(t *testing.T) {
		l, err := m.Role.Users(mctx, r.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", l.Users)
	})

	t.Run("AssociatePermissions", func(t *testing.T) {
		ps := []*Permission{
			{Name: auth0.String("read:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
			{Name: auth0.String("update:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
		}
		err = m.Role.AssociatePermissions(mctx, r.GetID(), ps...)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Permissions", func(t *testing.T) {
		l, err := m.Role.Permissions(mctx, r.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", l.Permissions)
	})

	t.Run("RemovePermissions", func(t *testing.T) {
		ps := []*Permission{
			{Name: auth0.String("read:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
			{Name: auth0.String("update:resource"), ResourceServerIdentifier: auth0.String("https://api.example.com/role")},
		}
		err = m.Role.RemovePermissions(mctx, r.GetID(), ps...)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Role.Delete(mctx, r.GetID())
		if err != nil {
			t.Error(err)
		}
	})
}
