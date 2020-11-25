package management

import (
	"encoding/json"
	"testing"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/internal/testing/expect"
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

		t.Run("BrandingColors", func(t *testing.T) {

			err = m.Branding.Update(&Branding{
				Colors: &BrandingColors{
					Primary: auth0.String("#ea5323"),
					PageBackgroundGradient: &BrandingPageBackgroundGradient{
						Type:        auth0.String("linear-gradient"),
						Start:       auth0.String("#000000"),
						End:         auth0.String("#ffffff"),
						AngleDegree: auth0.Int(35),
					},
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
	})
}

func TestBrandingColors(t *testing.T) {
	var tests = []struct {
		name   string
		colors *BrandingColors
		expect string
	}{
		{
			name: "PageBackground",
			colors: &BrandingColors{
				Primary:        auth0.String("#ea5323"),
				PageBackground: auth0.String("#000000"),
			},
			expect: `{"primary":"#ea5323","page_background":"#000000"}`,
		},
		{
			name: "PageBackgroundGradient",
			colors: &BrandingColors{
				Primary: auth0.String("#ea5323"),
				PageBackgroundGradient: &BrandingPageBackgroundGradient{
					Type:        auth0.String("linear-gradient"),
					Start:       auth0.String("#000000"),
					End:         auth0.String("#ffffff"),
					AngleDegree: auth0.Int(35),
				},
			},
			expect: `{"primary":"#ea5323","page_background":{"type":"linear-gradient","start":"#000000","end":"#ffffff","angle_deg":35}}`,
		},
		{
			name: "PageBackgroundNil",
			colors: &BrandingColors{
				Primary: auth0.String("#ea5323"),
			},
			expect: `{"primary":"#ea5323"}`,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			b, err := json.Marshal(tt.colors)
			if err != nil {
				t.Error(err)
			}

			expect.Expect(t, string(b), tt.expect)

			var colors BrandingColors
			err = json.Unmarshal([]byte(tt.expect), &colors)
			if err != nil {
				t.Error(err)
			}

			expect.Expect(t, &colors, tt.colors)
		})
	}
}

func TestBrandingTemplateUniversalLogin(t *testing.T) {
	var btul *BrandingTemplateUniversalLogin
	var err error
	var merr *managementError
	var ok bool

	m.Branding.DeleteTemplateUniversalLogin()

	t.Run("ReadTemplateUniversalLogin", func(t *testing.T) {
		btul, err = m.Branding.ReadTemplateUniversalLogin()
		if btul != nil {
			t.Fatalf("unexpected output. have %v, expected %v", btul, nil)
		}
		merr, ok = err.(*managementError)
		if !ok {
			t.Fatal(err)
		}
		expect.Expect(t, 404, merr.StatusCode)
	})

	t.Run("UpdateTemplateUniversalLogin", func(t *testing.T) {
		err = m.Branding.UpdateTemplateUniversalLogin(&BrandingTemplateUniversalLogin{
			Body: auth0.String("<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>"),
		})
		if err != nil {
			t.Fatal(err)
		}

		btul, err = m.Branding.ReadTemplateUniversalLogin()
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, *btul.Body, "<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>")
		t.Logf("%#v\n", btul)
	})

	t.Run("DeleteTemplateUniversalLogin", func(t *testing.T) {
		err = m.Branding.DeleteTemplateUniversalLogin()
		if err != nil {
			t.Fatal(err)
		}

		btul, err = m.Branding.ReadTemplateUniversalLogin()
		if btul != nil {
			t.Fatalf("unexpected output. have %v, expected %v", btul, nil)
		}
		merr, ok = err.(*managementError)
		if !ok {
			t.Fatal(err)
		}
		expect.Expect(t, 404, merr.StatusCode)
	})
}
