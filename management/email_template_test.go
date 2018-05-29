package management

import (
	"net/http"
	"testing"
)

func TestEmailTemplate(t *testing.T) {

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

	err := m.Email.Create(e)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Email.Delete()

	et := &EmailTemplate{
		Template:  "verify_email",
		Body:      "<html><body><h1>Verify your email</h1></body></html>",
		From:      "me@example.com",
		ResultURL: "https://www.example.com/verify-email",
		Subject:   "Verify your email",
		Syntax:    "liquid",
		Enabled:   true,
	}

	t.Run("Create", func(t *testing.T) {
		err = m.EmailTemplate.Create(et)
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() != http.StatusConflict {
				t.Fatal(err)
			}
		}
		t.Logf("%v\n", et)
	})

	t.Run("Read", func(t *testing.T) {
		et, err = m.EmailTemplate.Read(et.Template)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", et)
	})

	t.Run("Update", func(t *testing.T) {
		e.Enabled = true
		err = m.EmailTemplate.Update(et.Template, et)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", et)
	})

	t.Run("Delete", func(t *testing.T) {
		et.Enabled = false
		err = m.EmailTemplate.Update(et.Template, et)
		if err != nil {
			t.Error(err)
		}
	})
}
