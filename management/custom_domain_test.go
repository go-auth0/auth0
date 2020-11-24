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

	err = m.CustomDomain.Create(c)
	if err != nil {
		err, ok := err.(Error)
		if ok && err.Status() == http.StatusForbidden {
			// skip for free testing tenant
			t.Skip(err)
		} else if ok && err.Status() == http.StatusConflict {
			// only one custom domain available
			cs, err := m.CustomDomain.List()
			if err != nil {
				t.Fatal(err)
			}
			if len(cs) > 0 {
				m.CustomDomain.Delete(cs[0].GetID())
				defer func() {
					c := &CustomDomain{
						Domain:             cs[0].Domain,
						Type:               cs[0].Type,
						VerificationMethod: cs[0].VerificationMethod,
					}
					m.CustomDomain.Create(c)
					m.CustomDomain.Verify(c.GetID())
				}()
			}
		} else {
			t.Error(err)
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
		c, err = m.CustomDomain.Verify(c.GetID())
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
