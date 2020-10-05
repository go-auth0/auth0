package management

import (
	"testing"

	"gopkg.in/auth0.v5/internal/testing/expect"

	"gopkg.in/auth0.v5"
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
		err = m.Hook.Create(mctx, r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.Hook.Read(mctx, r.GetID())
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

		err = m.Hook.Update(mctx, id, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("List", func(t *testing.T) {
		r, err := m.Hook.List(mctx)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Hook.Delete(mctx, r.GetID())
		if err != nil {
			t.Error(err)
		}
	})
}

func TestHookSecrets(t *testing.T) {

	r := &HookSecrets{
		"SECRET1": "value1",
		"SECRET2": "value2",
	}

	hook := &Hook{
		Name:      auth0.String("test-hook-secrets"),
		Script:    auth0.String("function (user, context, callback) { callback(null, { user }); }"),
		TriggerID: auth0.String("pre-user-registration"),
		Enabled:   auth0.Bool(false),
	}

	err := m.Hook.Create(mctx, hook)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err = m.Hook.Delete(mctx, hook.GetID()); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create", func(t *testing.T) {
		err = m.Hook.CreateSecrets(mctx, hook.GetID(), r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Update", func(t *testing.T) {
		(*r)["SECRET1"] = "othervalue"
		delete(*r, "SECRET2") // patch allows only specifying one property
		err = m.Hook.UpdateSecrets(mctx, hook.GetID(), r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		result, err := m.Hook.Secrets(mctx, hook.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)

		expect.Expect(t, (*result)["SECRET1"], "_VALUE_NOT_SHOWN_")
		expect.Expect(t, (*result)["SECRET2"], "_VALUE_NOT_SHOWN_")
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Hook.RemoveSecrets(mctx, hook.GetID(), "SECRET1")
		if err != nil {
			t.Error(err)
		}

		result, err := m.Hook.Secrets(mctx, hook.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)

		expect.Expect(t, (*result)["SECRET1"], "")
		expect.Expect(t, (*result)["SECRET2"], "_VALUE_NOT_SHOWN_")
	})

	t.Run("RemoveAllSecrets", func(t *testing.T) {
		err = m.Hook.RemoveAllSecrets(mctx, hook.GetID())
		if err != nil {
			t.Error(err)
		}

		r = &HookSecrets{
			"SECRET3": "secret3",
		}

		err = m.Hook.CreateSecrets(mctx, hook.GetID(), r)
		if err != nil {
			t.Fatal(err)
		}

		result, err := m.Hook.Secrets(mctx, hook.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)

		expect.Expect(t, (*result)["SECRET1"], "")
		expect.Expect(t, (*result)["SECRET2"], "")
		expect.Expect(t, (*result)["SECRET3"], "_VALUE_NOT_SHOWN_")
	})
}
