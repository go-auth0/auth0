package management

import (
	"net/http"
	"testing"
	"time"

	"gopkg.in/auth0.v5"
)

func TestCustomDomain(t *testing.T) {

	c := &CustomDomain{
		Domain:             auth0.Stringf("auth.%d.alexkappa.com", time.Now().UTC().Unix()),
		Type:               auth0.String("auth0_managed_certs"),
		VerificationMethod: auth0.String("txt"),
	}

	var err error

	// setup
	c, err = m.CustomDomain.Read(c.GetID())
	if err != nil {
		if err, ok := err.(Error); ok && err.Status() == http.StatusNotFound {
			t.Skip(err)
		}
	} else {
		m.CustomDomain.Delete(c.GetID())
	}

	t.Run("Create", func(t *testing.T) {
		err = m.CustomDomain.Create(c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.CustomDomain.Read(c.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Verify", func(t *testing.T) {
		c, err := m.CustomDomain.Verify(c.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.CustomDomain.Delete(c.GetID())
		if err != nil {
			t.Error(err)
		}
	})
}
