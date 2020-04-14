package management

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/internal/testing/expect"
)

func TestClient(t *testing.T) {

	c := &Client{
		Name: auth0.Stringf("Test Client (%s)",
			time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client."),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Client.Create(c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.Client.Read(auth0.StringValue(c.ClientID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("List", func(t *testing.T) {
		var cl *ClientList
		cl, err = m.Client.List(WithFields("client_id"))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", cl)
	})

	t.Run("Update", func(t *testing.T) {
		id := auth0.StringValue(c.ClientID)

		c.ClientID = nil                       // read-only
		c.JWTConfiguration.SecretEncoded = nil // read-only
		c.SigningKeys = nil                    // read-only
		c.Description = auth0.String(strings.Replace(auth0.StringValue(c.Description), "just", "more than", 1))

		err = m.Client.Update(id, c)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", c)
	})

	t.Run("RotateSecret", func(t *testing.T) {
		secret := c.GetClientSecret()
		c, err = m.Client.RotateSecret(c.GetClientID())
		if err != nil {
			t.Error(err)
		}
		if secret == c.GetClientSecret() {
			t.Errorf("expected secret to change but didn't")
		}
		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Client.Delete(c.GetClientID())
		if err != nil {
			t.Error(err)
		}
	})
}

func TestClientJwtConfig(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		for name, tt := range map[string]struct {
			in       *ClientJWTConfiguration
			expected string
		}{
			"Nil":         {&ClientJWTConfiguration{Scopes: nil}, `{}`},
			"NilMap":      {&ClientJWTConfiguration{Scopes: (map[string]interface{})(nil)}, `{}`},
			"EmptyMap":    {&ClientJWTConfiguration{Scopes: map[string]interface{}{}}, `{"scopes":{}}`},
			"NonEmptyMap": {&ClientJWTConfiguration{Scopes: map[string]interface{}{"foo": "bar"}}, `{"scopes":{"foo":"bar"}}`},
		} {
			t.Run(name, func(t *testing.T) {
				b, err := json.Marshal(tt.in)
				if err != nil {
					t.Error(err)
				}
				expect.Expect(t, string(b), tt.expected)
			})
		}
	})
	t.Run("Unmarshal", func(t *testing.T) {
		for name, tt := range map[string]struct {
			in       string
			expected *ClientJWTConfiguration
		}{
			"Nil":         {`{}`, &ClientJWTConfiguration{Scopes: nil}},
			"EmptyMap":    {`{"scopes":{}}`, &ClientJWTConfiguration{Scopes: map[string]interface{}{}}},
			"NonEmptyMap": {`{"scopes":{"foo":"bar"}}`, &ClientJWTConfiguration{Scopes: map[string]interface{}{"foo": "bar"}}},
		} {
			t.Run(name, func(t *testing.T) {
				var out ClientJWTConfiguration
				if err := json.Unmarshal([]byte(tt.in), &out); err != nil {
					t.Error(err)
				}
				expect.Expect(t, &out, tt.expected)
			})
		}
	})
}
