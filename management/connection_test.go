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

func TestConnectionOptionsEmail_UnmarshalJSON(t *testing.T) {
	t.Run("Unmarshal passwordless email option", func(t *testing.T) {
		pe := PasswordlessEmail{
			Syntax:  auth0.String("liquid"),
			From:    auth0.String("from@example.com"),
			Subject: auth0.String("a subject"),
			Body:    auth0.String("<html><body>Some body</body></html>"),
		}
		oe := &ConnectionOptionsEmail{v: &pe}
		co := ConnectionOptions{
			Name:  auth0.String("email"),
			Email: oe,
		}
		s := Stringify(co)

		var aco ConnectionOptions
		err := json.Unmarshal([]byte(s), &aco)

		if err != nil {
			t.Error(err)
		}

		if ape, ok := aco.Email.PasswordlessEmail(); !ok {
			t.Fatal("ConnectionOptions.Email value is not of type *ConnectionOptionsEmail")
		} else {
			if *ape.Syntax != *pe.Syntax {
				t.Fatalf("expected syntax to be %v but got %v", *pe.Syntax, *ape.Syntax)
			}
			if *ape.From != *pe.From {
				t.Fatalf("expected from to be %v but got %v", *pe.From, *ape.From)
			}
			if *ape.Subject != *pe.Subject {
				t.Fatalf("expected subject to be %v but got %v", *pe.Subject, *ape.Subject)
			}
			if *ape.Body != *pe.Body {
				t.Fatalf("expected body to be %v but got %v", *pe.Body, *ape.Body)
			}
		}
	})

	t.Run("Unmarshal bool email option", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{v: auth0.Bool(true)}
		co := ConnectionOptions{Email: oe}
		s := Stringify(co)

		var actual ConnectionOptions
		err := json.Unmarshal([]byte(s), &actual)

		if err != nil {
			t.Error(err)
		}

		if e, ok := actual.Email.Bool(); !ok {
			t.Fatal("Email value is not of type *bool")
		} else {
			if !*e {
				t.Fatal("expected email to be true")
			}
		}
	})

	t.Run("Unmarshal unexpected email option", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{v: auth0.Time(time.Now())}
		co := ConnectionOptions{Email: oe}
		s := Stringify(co)

		var actual ConnectionOptions
		err := json.Unmarshal([]byte(s), &actual)

		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestConnectionOptionsEmail_MarshalJSON(t *testing.T) {
	type OptionsTest struct {
		Email *ConnectionOptionsEmail `json:"email,omitempty"`
	}

	t.Run("Marshal passwordless email values", func(t *testing.T) {
		pe := &PasswordlessEmail{
			Syntax:  auth0.String("syntax"),
			From:    auth0.String("from"),
			Subject: auth0.String("subject"),
			Body:    auth0.String("body"),
		}
		oe := &ConnectionOptionsEmail{}
		oe.SetPasswordlessEmail(pe)
		ot := OptionsTest{Email: oe}

		b, err := json.Marshal(ot)
		if err != nil {
			t.Error(err)
		}

		actual := string(b)
		expected := `{"email":{"syntax":"syntax","from":"from","subject":"subject","body":"body"}}`

		if actual != expected {
			t.Fatalf("expected json to be %v but got %v", expected, actual)
		}
	})

	t.Run("Marshal boolean values", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{}
		oe.SetBool(auth0.Bool(true))
		ot := OptionsTest{Email: oe}

		b, err := json.Marshal(ot)
		if err != nil {
			t.Error(err)
		}

		actual := string(b)
		expected := `{"email":true}`

		if actual != expected {
			t.Fatalf("expected json to be %v but got %v", expected, actual)
		}
	})
}

func TestConnectionOptionsEmail_PasswordlessEmail(t *testing.T) {
	t.Run("get passwordless email", func(t *testing.T) {
		pe := &PasswordlessEmail{
			Syntax:  auth0.String("syntax"),
			From:    auth0.String("from"),
			Subject: auth0.String("subject"),
			Body:    auth0.String("body"),
		}
		oe := &ConnectionOptionsEmail{v: pe}

		actual, b := oe.PasswordlessEmail()
		if !b {
			t.Error("unexpected type")
		}
		if actual != pe {
			t.Fatalf("expected passwordless email to be %v but got %v", actual, pe)
		}
	})

	t.Run("nil passwordless email", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{}

		_, ok := oe.PasswordlessEmail()
		if ok {
			t.Error("expected not ok")
		}
	})

	t.Run("not a passwordless email", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{v: auth0.Bool(true)}

		_, ok := oe.PasswordlessEmail()
		if ok {
			t.Error("expected not ok")
		}
	})
}

func TestConnectionOptionsEmail_Bool(t *testing.T) {
	t.Run("get bool", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{v: auth0.Bool(true)}

		actual, b := oe.Bool()
		if !b {
			t.Error("unexpected type")
		}
		if !*actual {
			t.Fatalf("expected bool to be true")
		}
	})

	t.Run("nil bool", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{}

		_, ok := oe.Bool()
		if ok {
			t.Error("expected not ok")
		}
	})

	t.Run("not bool", func(t *testing.T) {
		oe := &ConnectionOptionsEmail{v: &PasswordlessEmail{}}

		_, ok := oe.Bool()
		if ok {
			t.Error("expected not ok")
		}
	})
}
