package management

import (
	"testing"
	"time"

	"gopkg.in/auth0.v3"
	"gopkg.in/auth0.v3/internal/testing/expect"
)

func TestConnection(t *testing.T) {

	c := &Connection{
		Name:     auth0.Stringf("Test-Connection-%d", time.Now().Unix()),
		Strategy: auth0.String("auth0"),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Connection.Create(c)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := c.Options.(*ConnectionOptions); !ok {
			t.Errorf("unexpected options type %T", c.Options)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.Connection.Read(auth0.StringValue(c.ID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("List", func(t *testing.T) {
		cs, err := m.Connection.List()
		if err != nil {
			t.Error(err)
		}
		for _, c := range cs.Connections {
			var ok bool
			switch c.GetStrategy() {
			case "auth0":
				_, ok = c.Options.(*ConnectionOptions)
			case "google-oauth2":
				_, ok = c.Options.(*ConnectionOptionsGoogleOAuth2)
			case "facebook":
				_, ok = c.Options.(*ConnectionOptionsFacebook)
			case "apple":
				_, ok = c.Options.(*ConnectionOptionsApple)
			case "linkedin":
				_, ok = c.Options.(*ConnectionOptionsLinkedin)
			case "github":
				_, ok = c.Options.(*ConnectionOptionsGitHub)
			case "windowslive":
				_, ok = c.Options.(*ConnectionOptionsWindowsLive)
			case "salesforce":
				_, ok = c.Options.(*ConnectionOptionsSalesforce)
			case "email":
				_, ok = c.Options.(*ConnectionOptionsEmail)
			case "sms":
				_, ok = c.Options.(*ConnectionOptionsSMS)
			case "oidc":
				_, ok = c.Options.(*ConnectionOptionsOIDC)
			case "waad":
				_, ok = c.Options.(*ConnectionOptionsAzureAD)
			default:
				_, ok = c.Options.(map[string]interface{})
			}

			if !ok {
				t.Errorf("unexpected options type %T", c.Options)
			}

			t.Logf("%s %s %T\n", c.GetID(), c.GetName(), c.Options)
		}
	})

	t.Run("Update", func(t *testing.T) {

		id := auth0.StringValue(c.ID)

		c.ID = nil       // read-only
		c.Name = nil     // read-only
		c.Strategy = nil // read-only

		c.Options = &ConnectionOptions{

			BruteForceProtection: auth0.Bool(true),
			ImportMode:           auth0.Bool(false), // try some zero values
			DisableSignup:        auth0.Bool(true),
			RequiresUsername:     auth0.Bool(false),

			CustomScripts: map[string]interface{}{"get_user": "function( { return callback(null) }"},
			Configuration: map[string]interface{}{"foo": "bar"},
		}

		err = m.Connection.Update(id, c)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Connection.Delete(auth0.StringValue(c.ID))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ReadByName", func(t *testing.T) {
		cs, err := m.Connection.ReadByName("Username-Password-Authentication")
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", cs)
	})

	t.Run("GoogleOAuth2", func(t *testing.T) {
		g := &Connection{
			Name:     auth0.Stringf("Test-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("google-oauth2"),
			Options: &ConnectionOptionsGoogleOAuth2{
				AllowedAudiences: []interface{}{
					"example.com",
					"api.example.com",
				},
				Profile:  auth0.Bool(true),
				Calendar: auth0.Bool(true),
				Youtube:  auth0.Bool(false),
			},
		}

		defer m.Connection.Delete(g.GetID())

		err := m.Connection.Create(g)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := g.Options.(*ConnectionOptionsGoogleOAuth2)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}

		expect.Expect(t, o.GetProfile(), true)
		expect.Expect(t, o.GetCalendar(), true)
		expect.Expect(t, o.GetYoutube(), false)
		expect.Expect(t, o.Scopes(), []string{"email", "profile", "calendar"})

		t.Logf("%s\n", g)
	})
}
