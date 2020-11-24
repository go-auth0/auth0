package management

import (
	"encoding/json"
	"testing"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/internal/testing/expect"
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
			SessionLifetime:       auth0.Float64(720),
			IdleSessionLifetime:   auth0.Float64(72.2), // will be rounded off
		})
		if err != nil {
			t.Error(err)
		}
		tn, _ = m.Tenant.Read()
		t.Logf("%v\n", tn)
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		for tenant, expected := range map[*Tenant]string{
			{}:                                         `{}`,
			{SessionLifetime: auth0.Float64(1.2)}:      `{"session_lifetime":1}`,
			{SessionLifetime: auth0.Float64(1.19)}:     `{"session_lifetime":1}`,
			{SessionLifetime: auth0.Float64(1)}:        `{"session_lifetime":1}`,
			{SessionLifetime: auth0.Float64(720)}:      `{"session_lifetime":720}`,
			{IdleSessionLifetime: auth0.Float64(1)}:    `{"idle_session_lifetime":1}`,
			{IdleSessionLifetime: auth0.Float64(1.2)}:  `{"idle_session_lifetime":1}`,
			{SessionLifetime: auth0.Float64(0.25)}:     `{"session_lifetime_in_minutes":15}`,
			{SessionLifetime: auth0.Float64(0.5)}:      `{"session_lifetime_in_minutes":30}`,
			{SessionLifetime: auth0.Float64(0.99)}:     `{"session_lifetime_in_minutes":59}`,
			{IdleSessionLifetime: auth0.Float64(0.25)}: `{"idle_session_lifetime_in_minutes":15}`,
		} {
			b, err := json.Marshal(tenant)
			if err != nil {
				t.Error(err)
			}
			expect.Expect(t, string(b), expected)
			t.Logf("%v\n", tenant)
		}
	})
}
