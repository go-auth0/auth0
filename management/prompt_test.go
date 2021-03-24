package management

import (
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/internal/testing/expect"
	"testing"
)

func TestPrompt(t *testing.T) {
	t.Cleanup(func() {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "classic",
			IdentifierFirst:          auth0.Bool(false),
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	// update to the new identifier first experience
	err := m.Prompt.Update(&Prompt{
		UniversalLoginExperience: "new",
		IdentifierFirst:          auth0.Bool(true),
	})
	if err != nil {
		t.Error(err)
	}

	ps, err := m.Prompt.Read()
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, ps.UniversalLoginExperience, "new")
	expect.Expect(t, ps.IdentifierFirst, auth0.Bool(true))

	// update to the classic non identifier first experience
	err = m.Prompt.Update(&Prompt{
		UniversalLoginExperience: "classic",
		IdentifierFirst:          auth0.Bool(false),
	})
	if err != nil {
		t.Error(err)
	}

	ps, err = m.Prompt.Read()
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, ps.UniversalLoginExperience, "classic")
	expect.Expect(t, ps.IdentifierFirst, auth0.Bool(false))
}
