package management

import "testing"

func TestRule(t *testing.T) {

	r := &Rule{
		Name:    "test-rule",
		Script:  "function (user, context, callback) { callback(null, user, context); }",
		Enabled: false,
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Rule.Create(r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.Rule.Read(r.ID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Update", func(t *testing.T) {
		id := r.ID
		r.ID = "" // read-only
		r.Enabled = true

		err = m.Rule.Update(id, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Rule.Delete(r.ID)
		if err != nil {
			t.Error(err)
		}
	})
}
