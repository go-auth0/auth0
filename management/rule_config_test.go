package management

import "testing"

func TestRuleConfig(t *testing.T) {

	key := "foo"
	r := &RuleConfig{Value: "bar"}

	var err error

	t.Run("Upsert", func(t *testing.T) {
		err = m.RuleConfig.Upsert(key, r)
		if err != nil {
			t.Error(err)
		}
		if r.Key != key {
			t.Errorf("key mismatch %q != %q", r.Key, key)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.RuleConfig.Read(key)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Upsert", func(t *testing.T) {
		r.Key = ""
		r.Value = "baz"
		err = m.RuleConfig.Upsert(key, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.RuleConfig.Delete(key)
		if err != nil {
			t.Error(err)
		}
	})
}
