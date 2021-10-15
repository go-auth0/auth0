package management

import (
	"testing"

	"gopkg.in/auth0.v5/internal/testing/expect"
)

func TestCaptcha(t *testing.T) {
	t.Run("GetCaptchaSettings", func(t *testing.T) {
		settings, err := m.Captcha.GetCaptchaSettings()
		if err != nil {
			t.Error(err)
		}
		expect.Expect(t, settings, &CaptchaSettings{
			Selected: "auth0",
			Policy:   "off",
			Providers: &CaptchaProviders{
				Auth0:               &CaptchaProviderAuth0{},
				RecaptchaV2:         &CaptchaProviderRecaptchaV2{},
				RecaptchaEnterprise: &CaptchaProviderRecaptchaEnterprise{},
			},
		})
	})

	t.Run("SetCaptchaSettings", func(t *testing.T) {
		c := &CaptchaSettings{
			Selected: "recaptcha_v2",
			Policy:   "high_risk",
			Providers: &CaptchaProviders{
				RecaptchaV2: &CaptchaProviderRecaptchaV2{
					SiteKey: "foo",
					Secret:  "bar",
				},
			},
		}

		err := m.Captcha.SetCaptchaSettings(c)
		if err != nil {
			t.Error(err)
		}
	})

}
