package management

import (
	"fmt"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {

	c := &Connection{
		Name:     fmt.Sprintf("Test-Connection-%d", time.Now().Unix()),
		Strategy: "auth0",
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
		c, err = m.Connection.Read(c.ID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Update", func(t *testing.T) {
		id := c.ID
		c.ID = ""       // read-only
		c.Name = ""     // read-only
		c.Strategy = "" // read-only
		c.Options = &ConnectionOptions{
			CustomScripts: map[string]interface{}{"get_user": "function(email, callback) { return callback(null) }"},
			Configuration: map[string]interface{}{"foo": "bar"},
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
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Connection.Delete(c.ID)
		if err != nil {
			t.Error(err)
		}
	})
}
