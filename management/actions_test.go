// +build integration

package management

import (
	"testing"
)

func TestActions(t *testing.T) {
	a := &Action{
		Name: "my-action-8",
		SupportedTriggers: []Trigger{
			{ID: PostLogin, Version: "v1"},
		},
	}

	v := &ActionVersion{
		Runtime: "node12",
		Code:    `module.exports = function(user, context, cb) { cb(null, user, context) }`,
		Dependencies: []Dependency{
			{Name: "lodash", Version: "v4.17.20"},
		},
	}

	t.Run("Create", func(t *testing.T) {
		err := m.Action.Create(a)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", a)
	})

	t.Run("Read", func(t *testing.T) {
		got, err := m.Action.Read(a.ID)
		if err != nil {
			t.Error(err)
		}
		if a.Name != got.Name {
			t.Fatalf("wanted name: %s, got %s", a.Name, got.Name)
		}
		t.Logf("=====> orig: %v\n", a)
		t.Logf("=====> got: %v\n", got)
	})

	t.Run("Update", func(t *testing.T) {
		a.Name = a.Name + "-renamed"

		// Typically the patch is only on name.
		a.SupportedTriggers = nil
		err := m.Action.Update(a.ID, a)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", a)
	})

	// t.Run("List", func(t *testing.T) {
	// 	list, err := m.Action.List(WithTriggerID(PostLogin))
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	t.Logf("%v\n", list.Actions)

	// })

	t.Run("Create Version", func(t *testing.T) {
		if err := m.ActionVersion.Create(a.ID, v); err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", v)
	})

	t.Run("Read Version", func(t *testing.T) {
		gotVersion, err := m.ActionVersion.Read(a.ID, v.ID)
		if err != nil {
			t.Error(err)
		}

		if want, got := v.ID, gotVersion.ID; want != got {
			t.Errorf("wanted ID: %s, got %s", want, got)
		}
		t.Logf("%v\n", v)
	})

	t.Run("Delete Version", func(t *testing.T) {
		if err := m.ActionVersion.Delete(a.ID, v.ID); err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", v)
	})

	// t.Run("Deploy + Test", func(t *testing.T) {
	// 	v := &ActionVersion{
	// 		Runtime: "node12",
	// 		Code:    `module.exports = function(user, context, cb) { cb(null, user, context) }`,
	// 		Dependencies: []Dependency{
	// 			{Name: "lodash", Version: "v4.17.20"},
	// 		},
	// 	}

	// 	if err := m.ActionVersion.Deploy(a.ID, v); err != nil {
	// 		t.Error(err)
	// 	}

	// 	t.Logf("%v\\n", v)

	// 	testPayload := Object{"user": struct{}{}, "context": struct{}{}}
	// 	resultPayload, err := m.ActionVersion.Test(a.ID, v.ID, testPayload)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	t.Logf("%v\\n", resultPayload)
	// })

	t.Run("Delete", func(t *testing.T) {
		err := m.Action.Delete(a.ID)
		if err != nil {
			t.Error(err)
		}
	})
}
