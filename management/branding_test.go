package management

import (
	"gopkg.in/auth0.v1"
	"testing"
)

func TestBranding(t *testing.T) {

	var branding *Branding
	var err error

	t.Run("Read", func(t *testing.T) {
		branding, err = m.Branding.Read()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", branding)
	})

	t.Run("Update", func(t *testing.T) {
		err = m.Branding.Update(&Branding{
			Colors: &BrandingColors{
				Primary:        auth0.String("#ea5323"),
				PageBackground: auth0.String("#000000"),
			},
			FaviconURL: auth0.String("https://mycompany.org/favicon.ico"),
			LogoURL:    auth0.String("https://mycompany.org/logo.png"),
			Font: &BrandingFont{
				URL: auth0.String("https://mycompany.org/font.otf"),
			},
		})
		if err != nil {
			t.Error(err)
		}
		branding, _ = m.Branding.Read()
		t.Logf("%v\n", branding)
	})
}
