package management

import (
	"testing"

	auth0 "github.com/yieldr/go-auth0"
)

func TestChangePassword(t *testing.T) {

	changePw := &DBConnectionsChangePassword{
		ClientID:   auth0.String("Auth0 clientID"),
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
		Password:   auth0.String("Passwords hide their Chuck"),
	}

	t.Run("ChangePassword", func(t *testing.T) {
		str, err := m.DBConnections.ChangePassword(changePw)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", str)
	})
}
