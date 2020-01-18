package management

import (
	"testing"
	"time"

	"gopkg.in/auth0.v3"
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
		t.Logf("%v\n", cs)
	})

	t.Run("Update", func(t *testing.T) {

		id := auth0.StringValue(c.ID)

		c.ID = nil       // read-only
		c.Name = nil     // read-only
		c.Strategy = nil // read-only

		c.Options = &ConnectionOptions{
			ExtAdmin:       auth0.Bool(true),
			ExtGroups:      auth0.Bool(true),
			ExtProfile:     auth0.Bool(true),
			ExtIsSuspended: auth0.Bool(false), // try some zero values
			ExtAgreedTerms: auth0.Bool(false),

			CustomScripts: map[string]interface{}{"get_user": "function( { return callback(null) }"},
			Configuration: map[string]interface{}{"foo": "bar"},

			RequiresUsername: auth0.Bool(true),
		}

		cc := c // make a copy so we can compare later

		err = m.Connection.Update(id, c)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", c)

		if c.Options.CustomScripts["get_user"] != cc.Options.CustomScripts["get_user"] {
			t.Fatal(`unexpected result for "get_user" custom script`)
		}

		if _, exist := c.Options.Configuration["foo"]; !exist {
			t.Fatal(`missing key "foo"`)
		}

		if c.Options.RequiresUsername != cc.Options.RequiresUsername {
			t.Fatalf("expected requires_username to be %v but got %v", cc.Options.RequiresUsername, c.Options.RequiresUsername)
		}
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
}
