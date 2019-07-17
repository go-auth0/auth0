package management

import "encoding/json"

type Branding struct {
	// Change password page settings
	Colors *BrandingColors `json:"colors,omitempty"`

	// URL for the favicon. Must use HTTPS.
	FaviconURL *string `json:"favicon_url,omitempty"`

	// URL for the logo. Must use HTTPS.
	LogoURL *string `json:"logo_url,omitempty"`

	Font *BrandingFont `json:"font,omitempty"`
}

func (branding *Branding) String() string {
	b, _ := json.MarshalIndent(branding, "", "  ")
	return string(b)
}

type BrandingColors struct {
	// Accent color
	Primary *string `json:"primary,omitempty"`
	// Page background color
	PageBackground *string `json:"page_background,omitempty"`
}

type BrandingFont struct {
	// URL for the custom font. Must use HTTPS.
	URL *string `json:"url,omitempty"`
}

type BrandingManager struct {
	m *Management
}

func NewBrandingManager(m *Management) *BrandingManager {
	return &BrandingManager{m}
}

func (bm *BrandingManager) Read(opts ...reqOption) (*Branding, error) {
	branding := new(Branding)
	err := bm.m.get(bm.m.uri("branding")+bm.m.q(opts), branding)
	return branding, err
}

func (bm *BrandingManager) Update(t *Branding) (err error) {
	return bm.m.patch(bm.m.uri("branding"), t)
}
