package management

import (
	"testing"

	"gopkg.in/auth0.v5"
)

func TestTicket(t *testing.T) {

	var err error

	c, err := m.Connection.ReadByName("Username-Password-Authentication")
	if err != nil {
		t.Error(err)
	}

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@chucknorris.com"),
		Password:   auth0.String("I have a password and its a secret"),
	}
	if auth0.BoolValue(c.Options.(*ConnectionOptions).RequiresUsername) {
		u.Username = auth0.String("Chuck")
	}
	m.User.Create(u)

	userID := auth0.StringValue(u.ID)

	defer m.User.Delete(userID)

	t.Run("VerifyEmail", func(t *testing.T) {

		v := &Ticket{
			ResultURL: auth0.String("https://example.com/verify-email"),
			UserID:    auth0.String(userID),
			TTLSec:    auth0.Int(3600),
		}

		err = m.Ticket.VerifyEmail(v)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", v)
	})

	t.Run("ChangePassword", func(t *testing.T) {

		v := &Ticket{
			ResultURL:           auth0.String("https://example.com/change-password"),
			UserID:              auth0.String(userID),
			TTLSec:              auth0.Int(3600),
			MarkEmailAsVerified: auth0.Bool(true),
		}

		err = m.Ticket.ChangePassword(v)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", v)
	})
}
