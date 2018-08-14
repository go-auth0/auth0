package management

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	c := &Client{
		Name: fmt.Sprintf("Test Client (%s)",
			time.Now().Format(time.StampMilli)),
		Description: "This is just a test client.",
		Callbacks: []interface{}{"https://example.com/saml"},
		Addons: &Addons{
			Samlp: &Samlp{
				Audience: "https://example.com/saml",
				Mappings: &SamlMappings{
					Email: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
					Name: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name",
				},
				CreateUpnClaim: false,
				PassthroughClaimsWithNoMapping: false,
				MapUnknownClaimsAsIs: false,
				MapIdentities: false,
				NameIdentifierFormat: "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
				NameIdentifierProbes: []string{
					"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
				},
			},
		},
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
		c, err = m.Client.Read(c.ClientID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("List", func(t *testing.T) {
		var cs []*Client
		cs, err = m.Client.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", cs)
	})

	t.Run("Update", func(t *testing.T) {
		id := c.ClientID
		c.ClientID = "" // read-only
		c.Description = strings.Replace(c.Description, "just", "more than", 1)
		err = m.Client.Update(id, c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("RotateSecret", func(t *testing.T) {
		id := c.ClientID
		secret := c.ClientSecret
		c, err = m.Client.RotateSecret(id)
		if err != nil {
			t.Error(err)
		}
		if secret == c.ClientSecret {
			t.Errorf("expected secret to change but didn't")
		}
		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Client.Delete(c.ClientID)
		if err != nil {
			t.Error(err)
		}
	})
}
