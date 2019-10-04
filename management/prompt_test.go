package management

import (
	"testing"
)

func TestPrompt(t *testing.T) {
	defer func() {
		m.Prompt.UpdateSettings(&PromptSettings{
			UniversalLoginExperience: "classic",
		})
	}()

	t.Run("GetSettings", func(t *testing.T) {
		ps, err := m.Prompt.GetSettings()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", ps)
	})

	t.Run("UpdateSettings", func(t *testing.T) {
		expected := "new"
		err := m.Prompt.UpdateSettings(&PromptSettings{
			UniversalLoginExperience: expected,
		})
		if err != nil {
			t.Error(err)
		}
		ps, err := m.Prompt.GetSettings()
		if err != nil {
			t.Error(err)
		}
		if ps.UniversalLoginExperience != "new" {
			t.Errorf("unexpected output. have %v, expected %v", ps.UniversalLoginExperience, expected)
		}
		t.Logf("%v\n", ps)
	})
}
