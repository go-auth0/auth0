package management

import (
	"gopkg.in/auth0.v5"
	"encoding/json"
	"gopkg.in/auth0.v5/internal/testing/expect"

	"testing"

	"gopkg.in/auth0.v5/internal/testing/expect"
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

func TestPromptCustomText(t *testing.T) {

	t.Run("ReadCustomText", func(t *testing.T) {
		prompts := [16]string{
			PromptConsent,
			PromptDeviceFlow,
			PromptEmailOtpChallengeFlow,
			PromptEmailVerificationFlow,
			PromptLogin,
			PromptLoginEmailVerification,
			PromptMfa,
			PromptMfaEmail,
			PromptMfaOtp,
			PromptMfaPhone,
			PromptMfaPush,
			PromptMfaRecoveryCode,
			PromptMfaSms,
			PromptMfaVoice,
			PromptResetPassword,
			PromptSignup,
		}
		for _, prompt := range prompts {
			pct, err := m.Prompt.ReadCustomText(prompt, "en")
			if err != nil {
				t.Error(err)
			}
			expect.Expect(t, pct.Prompt, prompt)
			expect.Expect(t, pct.Language, "en")
			b, err := json.Marshal(pct.Screens)
			if err != nil {
				t.Error(err)
			}
			expect.Expect(t, string(b), "{}")
			t.Logf("%v\n", pct)
		}
	})

	t.Run("UpdateCustomText", func(t *testing.T) {
		defer m.Prompt.UpdateCustomText(&PromptCustomText{
			Prompt: PromptConsent,
			Language: "en",
			Screens: &ConsentScreens{},
		})

		err := m.Prompt.UpdateCustomText(&PromptCustomText{
			Prompt: PromptConsent,
			Language: "en",
			Screens: &ConsentScreens{
				Consent: map[string]interface{}{ "pageTitle": "new page title" },
			},
		})
		if err != nil {
			t.Error(err)
		}
		pct, err := m.Prompt.ReadCustomText(PromptConsent, "en")
		if err != nil {
			t.Error(err)
		}
		expect.Expect(t, pct.Prompt, PromptConsent)
		expect.Expect(t, pct.Language, "en")
		b, err := json.Marshal(pct.Screens)
		if err != nil {
			t.Error(err)
		}
		expect.Expect(t, string(b), `{"consent":{"pageTitle":"new page title"}}`)
		t.Logf("%v\n", pct)
	})
}