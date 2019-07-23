package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitBranding(t *testing.T) {

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
		wanted := `{}`
		if brandingInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", brandingInJson, wanted)
		}
	})
}
