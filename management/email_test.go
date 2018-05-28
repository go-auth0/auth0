package management

import "testing"

func TestEmail(t *testing.T) {

	e := &Email{
		Name:               "smtp",
		Enabled:            true,
		DefaultFromAddress: "no-reply@example.com",
		Credentials: &EmailCredentials{
			SMTPHost: "smtp.example.com",
			SMTPPort: 587,
			SMTPUser: "user",
			SMTPPass: "pass",
		},
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Email.Create(e)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", e)
	})

	t.Run("Read", func(t *testing.T) {
		e, err = m.Email.Read()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", e)
	})

	t.Run("Update", func(t *testing.T) {
		e.Enabled = true
		err = m.Email.Update(e)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", e)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Email.Delete()
		if err != nil {
			t.Error(err)
		}
	})
}
