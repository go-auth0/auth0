package management

import (
	"encoding/json"
	"gopkg.in/auth0.v1"
	"strings"
	"testing"
)

func TestBranding(t *testing.T) {

	var branding *Branding
	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		brandingInJson := `{
			"colors": {
				"primary": "#ea5323",
				"page_background": "#000000"
			},
			"favicon_url": "https://mycompany.org/favicon.ico",
			"logo_url": "https://mycompany.org/logo.png",
			"font": {
				"url": "https://mycompany.org/font.otf"
			}
		}`

		decoder := json.NewDecoder(strings.NewReader(brandingInJson))
		decoder.DisallowUnknownFields()
		var branding Branding
		err = decoder.Decode(&branding)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		branding := Branding{}
		bytes, _ := json.Marshal(branding)
		brandingInJson := string(bytes)
		wanted := "{}"
		if brandingInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", brandingInJson, wanted)
		}
	})

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
