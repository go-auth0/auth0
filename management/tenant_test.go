package management

import (
	"testing"

	"gopkg.in/auth0.v3"
)

func TestTenant(t *testing.T) {

	var tn *Tenant
	var err error

	t.Run("Read", func(t *testing.T) {
		tn, err = m.Tenant.Read()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", tn)
	})

	t.Run("Update", func(t *testing.T) {
		err = m.Tenant.Update(&Tenant{
			FriendlyName:          auth0.String("My Example Tenant"),
			SupportURL:            auth0.String("https://support.example.com"),
			SupportEmail:          auth0.String("support@example.com"),
			DefaultRedirectionURI: auth0.String("https://example.com/login"),
		})
		if err != nil {
			t.Error(err)
		}
		tn, _ = m.Tenant.Read()
		t.Logf("%v\n", tn)
	})
}
