package management

import (
	"testing"

	auth0 "github.com/yieldr/go-auth0"
)

func TestChangePassword(t *testing.T) {

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
		Password:   auth0.String("5301111b-b31b-47c4-bf3d-0c26ea57bdf4"),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.User.Create(u)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", u)
	})

	changePw := &DBConnectionsChangePassword{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
	}

	t.Run("ChangePassword", func(t *testing.T) {
		str, err := m.DBConnections.ChangePassword(changePw)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", str)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.User.Delete(auth0.StringValue(u.ID))
		if err != nil {
			t.Error(err)
		}
	})
}
