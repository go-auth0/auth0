package management

import (
	"encoding/json"
	"reflect"
	"testing"

	"gopkg.in/auth0.v4"
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

	t.Run("Update with gradient", func(t *testing.T) {
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
}

func TestBrandingColors(t *testing.T) {
	var serializerTests = []struct {
		name string
		bc   *BrandingColors
		json string
	}{
		{
			name: "Solid background",
			bc: &BrandingColors{
				Primary:        auth0.String("#ea5323"),
				PageBackground: auth0.String("#000000"),
			},
			json: "{\"primary\":\"#ea5323\",\"page_background\":\"#000000\"}",
		},
		{
			name: "Gradient background",
			bc: &BrandingColors{
				Primary: auth0.String("#ea5323"),
				PageBackgroundGradient: &BrandingPageBackgroundGradient{
					Type:        auth0.String("linear-gradient"),
					Start:       auth0.String("#000000"),
					End:         auth0.String("#ffffff"),
					AngleDegree: auth0.Int(35),
				},
			},
			json: "{\"primary\":\"#ea5323\",\"page_background\":{\"type\":\"linear-gradient\",\"start\":\"#000000\",\"end\":\"#ffffff\",\"angle_deg\":35}}",
		},
		{
			name: "No background",
			bc: &BrandingColors{
				Primary: auth0.String("#ea5323"),
			},
			json: "{\"primary\":\"#ea5323\"}",
		},
	}

	for _, tt := range serializerTests {
		t.Run(tt.name, func(t *testing.T) {
			// Check custom Marshal
			serializedBrandingColors, err := json.Marshal(tt.bc)
			if err != nil {
				t.Error(err)
			}

			if string(serializedBrandingColors) != tt.json {
				t.Errorf("serialization: expected %v, got %v", tt.json, string(serializedBrandingColors))
			}

			// Check custom Unmarshal
			var deserializedBrandingColors BrandingColors
			err = json.Unmarshal([]byte(tt.json), &deserializedBrandingColors)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(&deserializedBrandingColors, tt.bc) {
				t.Errorf("deserialization: expected %v, got %v", tt.bc, &deserializedBrandingColors)
			}
		})
	}
}
