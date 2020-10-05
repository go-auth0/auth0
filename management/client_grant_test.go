package management

import (
	"testing"
	"time"

	"gopkg.in/auth0.v5"
)

func TestClientGrant(t *testing.T) {

	var err error

	// We need a client and resource server to connect using a client grant. So
	// first we must create them.

	c := &Client{
		Name: auth0.Stringf("Test Client - Client Grant (%s)",
			time.Now().Format(time.StampMilli)),
	}

	err = m.Client.Create(mctx, c)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Client.Delete(mctx, auth0.StringValue(c.ClientID))

	s := &ResourceServer{
		Name: auth0.Stringf("Test Client Grant (%s)",
			time.Now().Format(time.StampMilli)),
		Identifier: auth0.String("https://api.example.com/client-grant"),
		Scopes: []*ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
			{
				Value:       auth0.String("update:resource"),
				Description: auth0.String("Update Resource"),
			},
		},
	}
	err = m.ResourceServer.Create(mctx, s)
	if err != nil {
		t.Fatal(err)
	}
	defer m.ResourceServer.Delete(mctx, auth0.StringValue(s.ID))

	g := &ClientGrant{
		ClientID: c.ClientID,
		Audience: s.Identifier,
		Scope:    []interface{}{"create:resource"},
	}

	t.Run("Create", func(t *testing.T) {
		err = m.ClientGrant.Create(mctx, g)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", g)
	})

	t.Run("Read", func(t *testing.T) {
		g, err = m.ClientGrant.Read(mctx, auth0.StringValue(g.ID))
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", g)
	})

	t.Run("Update", func(t *testing.T) {
		id := auth0.StringValue(g.ID)
		g.ID = nil
		g.Audience = nil // read-only
		g.ClientID = nil // read-only
		g.Scope = append(g.Scope, "update:resource")

		err = m.ClientGrant.Update(mctx, id, g)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", g)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.ClientGrant.Delete(mctx, auth0.StringValue(g.ID))
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("List", func(t *testing.T) {
		gs, err := m.ClientGrant.List(
			mctx,
			PerPage(10),          // overwrites the default 50
			IncludeTotals(false), // has no effect as it is enforced internally
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", gs)
	})
}
