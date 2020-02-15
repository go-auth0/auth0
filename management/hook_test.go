package management

import (
	"testing"

	"gopkg.in/auth0.v3"
)

func TestHook(t *testing.T) {

	r := &Hook{
		Name:      auth0.String("test-hook"),
		Script:    auth0.String("function (user, context, callback) { callback(null, { user }); }"),
		TriggerID: auth0.String("pre-user-registration"),
		Enabled:   auth0.Bool(false),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Hook.Create(r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.Hook.Read(r.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Update", func(t *testing.T) {
		id := r.GetID()

		r.ID = nil        // read-only
		r.TriggerID = nil // read-only
		r.Script = auth0.String("function (user, context, callback) { console.log('hooked!'); callback(null, { user }); }")
		r.Enabled = auth0.Bool(true)

		err = m.Hook.Update(id, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("List", func(t *testing.T) {
		r, err := m.Hook.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Hook.Delete(auth0.StringValue(r.ID))
		if err != nil {
			t.Error(err)
		}
	})
}
