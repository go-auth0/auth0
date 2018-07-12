package management

import "testing"

func TestTicket(t *testing.T) {

	var err error

	u := &User{
		Connection: "Username-Password-Authentication",
		Email:      "chuck@chucknorris.com",
		Password:   "I have a password and its a secret",
	}
	m.User.Create(u)
	defer m.User.Delete(u.ID)

	t.Run("VerifyEmail", func(t *testing.T) {

		v := &Ticket{
			ResultURL: "https://example.com/verify-email",
			UserID:    u.ID,
			TTLSec:    3600,
		}

		v, err = m.Ticket.VerifyEmail(v)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%#v\n", v)
	})

	t.Run("ChangePassword", func(t *testing.T) {

		v := &Ticket{
			ResultURL: "https://example.com/verify-email",
			UserID:    u.ID,
			TTLSec:    3600,
		}

		v, err = m.Ticket.ChangePassword(v)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%#v\n", v)
	})
}
