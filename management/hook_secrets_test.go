package management

import (
	"gopkg.in/auth0.v4/internal/testing/expect"
	"testing"

	"gopkg.in/auth0.v4"
)

func TestHookSecrets(t *testing.T) {

	r := &map[string]string{
		"SECRET1": "value1",
		"SECRET2": "value2",
	}

	hook := &Hook{
		Name:      auth0.String("test-hook-secrets"),
		Script:    auth0.String("function (user, context, callback) { callback(null, { user }); }"),
		TriggerID: auth0.String("pre-user-registration"),
		Enabled:   auth0.Bool(false),
	}

	err := m.Hook.Create(hook)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err = m.Hook.Delete(hook.GetID()); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Upsert", func(t *testing.T) {
		err = m.HookSecrets.Upsert(hook.GetID(), r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		result, err := m.HookSecrets.Read(hook.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)

		expect.Expect(t, (*result)["SECRET1"], "_VALUE_NOT_SHOWN_")
		expect.Expect(t, (*result)["SECRET2"], "_VALUE_NOT_SHOWN_")
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.HookSecrets.Delete(hook.GetID(), "SECRET1")
		if err != nil {
			t.Error(err)
		}

		result, err := m.HookSecrets.Read(hook.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)

		expect.Expect(t, (*result)["SECRET1"], "")
		expect.Expect(t, (*result)["SECRET2"], "_VALUE_NOT_SHOWN_")
	})
}
