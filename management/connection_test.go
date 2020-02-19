package management

import (
	"encoding/json"
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

func TestConnectionOptions_UnmarshalJSON(t *testing.T) {
	t.Run("Unmarshal passwordless email option", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{
			Syntax:  auth0.String("liquid"),
			From:    auth0.String("from@example.com"),
			Subject: auth0.String("a subject"),
			Body:    auth0.String("<html><body>Some body</body></html>"),
		}
		co := ConnectionOptions{
			Name:  auth0.String("email"),
			Email: oe,
		}
		s := Stringify(co)

		var actual ConnectionOptions
		err := json.Unmarshal([]byte(s), &actual)

		if err != nil {
			t.Error(err)
		}

		if e, ok := actual.Email.(*ConnectionOptionsEmail); !ok {
			t.Fatal("ConnectionOptions.Email value is not of type *ConnectionOptionsEmail")
		} else {
			if *e.Syntax != *oe.Syntax {
				t.Fatalf("expected syntax to be %v but got %v", *oe.Syntax, *e.Syntax)
			}
			if *e.From != *oe.From {
				t.Fatalf("expected from to be %v but got %v", *oe.From, *e.From)
			}
			if *e.Subject != *oe.Subject {
				t.Fatalf("expected subject to be %v but got %v", *oe.Subject, *e.Subject)
			}
			if *e.Body != *oe.Body {
				t.Fatalf("expected body to be %v but got %v", *oe.Body, *e.Body)
			}
		}
	})

	t.Run("Unmarshal unexpected email option", func(t *testing.T) {
		co := ConnectionOptions{
			Email: auth0.Bool(true),
		}
		s := Stringify(co)

		var actual ConnectionOptions
		err := json.Unmarshal([]byte(s), &actual)

		if err != nil {
			t.Error(err)
		}

		if e, ok := actual.Email.(*bool); !ok {
			t.Fatal("Email value is not of type *bool")
		} else {
			if !*e {
				t.Fatal("expected email to be true")
			}
		}
	})

	t.Run("Unmarshal unexpected email option", func(t *testing.T) {
		co := ConnectionOptions{
			Email: auth0.Time(time.Now()),
		}
		s := Stringify(co)

		var actual ConnectionOptions
		err := json.Unmarshal([]byte(s), &actual)

		if err == nil {
			t.Error("expected an error")
		}
	})
}
