package management

import (
	"net/http"
	"testing"
)

func TestCustomDomain(t *testing.T) {

	c := &CustomDomain{
		Domain:             "auth.example.com",
		Type:               "auth0_managed_certs",
		VerificationMethod: "txt",
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.CustomDomain.Create(c)
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() == http.StatusForbidden {
				t.Skip(err)
			} else {
				t.Error(err)
			}
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.CustomDomain.Read(c.ID)
		if err != nil {
			if err, ok := err.(Error); ok && err.Status() == http.StatusNotFound {
				t.Skip(err)
			} else {
				t.Error(err)
			}
		}
		t.Logf("%v\n", c)
	})

	t.Run("Update", func(t *testing.T) {
		id := c.ID
		c.ID = "" // read-only
		c.Domain = "id.example.com"
		err = m.CustomDomain.Update(id, c)
		if err, ok := err.(Error); ok && err.Status() == http.StatusNotFound {
			t.Skip(err)
		} else {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.CustomDomain.Delete(c.ID)
		if err, ok := err.(Error); ok && err.Status() == http.StatusNotFound {
			t.Skip(err)
		} else {
			t.Error(err)
		}
	})
}
