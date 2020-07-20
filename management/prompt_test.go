package management

import (
	"testing"
)

func TestPrompt(t *testing.T) {

	t.Run("Read", func(t *testing.T) {
		ps, err := m.Prompt.Read()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", ps)
	})

	t.Run("Update", func(t *testing.T) {
		defer func() {
			m.Prompt.Update(&Prompt{
				UniversalLoginExperience: "classic",
			})
		}()
		expected := "new"
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: expected,
		})
		if err != nil {
			t.Error(err)
		}

		ps, err := m.Prompt.Read()
		if err != nil {
			t.Error(err)
		}
		if ps.UniversalLoginExperience != expected {
			t.Errorf("unexpected output. have %v, expected %v", ps.UniversalLoginExperience, expected)
		}
		t.Logf("%v\n", ps)
	})
}
