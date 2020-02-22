package management

import (
	"encoding/json"
)

type Branding struct {
	// Change password page settings
	Colors *BrandingColors `json:"colors,omitempty"`

	// URL for the favicon. Must use HTTPS.
	FaviconURL *string `json:"favicon_url,omitempty"`

	// URL for the logo. Must use HTTPS.
	LogoURL *string `json:"logo_url,omitempty"`

	Font *BrandingFont `json:"font,omitempty"`
}

type BrandingColors struct {
	// Accent color
	Primary *string `json:"primary,omitempty"`

	// Page background color
	// Only one of PageBackground and PageBackgroundGradient should be set.
	// If both fields are set, PageBackground takes priority.
	PageBackground *string `json:"-"`

	// Page background gradient
	// Only one of PageBackground and PageBackgroundGradient should be set.
	// If both fields are set, PageBackground takes priority.
	PageBackgroundGradient *BrandingPageBackgroundGradient `json:"-"`
}

type BrandingPageBackgroundGradient struct {
	Type        *string `json:"type,omitempty"`
	Start       *string `json:"start,omitempty"`
	End         *string `json:"end,omitempty"`
	AngleDegree *int    `json:"angle_deg,omitempty"`
}

type jsonPageBackgroundSolid struct {
	PageBackground *string `json:"page_background,omitempty"`
}

type jsonPageBackgroundGradient struct {
	PageBackgroundGradient *BrandingPageBackgroundGradient `json:"page_background,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
//
// It is required to handle the json field page_background, which can either
// be a hex color string, or an object describing a gradient.
func (bc *BrandingColors) MarshalJSON() ([]byte, error) {
	var data interface{} = struct {
		Primary *string `json:"primary,omitempty"`
	}{
		Primary: bc.Primary,
	}

	if bc.PageBackground != nil {
		data = struct {
			Primary *string `json:"primary,omitempty"`
			*jsonPageBackgroundSolid
		}{
			Primary: bc.Primary,
			jsonPageBackgroundSolid: &jsonPageBackgroundSolid{
				PageBackground: bc.PageBackground,
			},
		}
	} else if bc.PageBackgroundGradient != nil {
		data = struct {
			Primary *string `json:"primary,omitempty"`
			*jsonPageBackgroundGradient
		}{
			Primary: bc.Primary,
			jsonPageBackgroundGradient: &jsonPageBackgroundGradient{
				PageBackgroundGradient: bc.PageBackgroundGradient,
			},
		}
	}

	return json.Marshal(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle the json field page_background, which can either
// be a hex color string, or an object describing a gradient.
func (bc *BrandingColors) UnmarshalJSON(data []byte) error {
	// Unmarshal common fields
	type JSONBrandingColors BrandingColors

	var commonBrandingColors JSONBrandingColors

	err := json.Unmarshal(data, &commonBrandingColors)
	if err != nil {
		return err
	}

	bc.Primary = commonBrandingColors.Primary

	// Unmarshal page_background
	// Can be either a string with solid color
	var solidBrandingColor jsonPageBackgroundSolid
	err = json.Unmarshal(data, &solidBrandingColor)
	if err == nil {
		bc.PageBackground = solidBrandingColor.PageBackground
		bc.PageBackgroundGradient = nil
	}

	// Or object describing a color-gradient
	var gradientBrandingColor jsonPageBackgroundGradient
	err = json.Unmarshal(data, &gradientBrandingColor)
	if err == nil {
		bc.PageBackgroundGradient = gradientBrandingColor.PageBackgroundGradient
		bc.PageBackground = nil
	}

	return nil
}

type BrandingFont struct {
	// URL for the custom font. Must use HTTPS.
	URL *string `json:"url,omitempty"`
}

type BrandingManager struct {
	*Management
}

func newBrandingManager(m *Management) *BrandingManager {
	return &BrandingManager{m}
}

// Retrieve various settings related to branding.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get_branding
func (m *BrandingManager) Read() (*Branding, error) {
	branding := new(Branding)
	err := m.get(m.uri("branding"), branding)
	return branding, err
}

// Update various fields related to branding.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/patch_branding
func (m *BrandingManager) Update(t *Branding) (err error) {
	return m.patch(m.uri("branding"), t)
}
