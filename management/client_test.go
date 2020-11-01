package management

import (
	"strings"
	"testing"
	"time"

	"gopkg.in/auth0.v5"
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
