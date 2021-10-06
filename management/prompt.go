package management

import "gopkg.in/auth0.v5"

type Prompt struct {
	// Which login experience to use. Can be `new` or `classic`.
	UniversalLoginExperience string `json:"universal_login_experience,omitempty"`

	// IdentifierFirst determines if the login screen prompts for just the identifier, identifier and password first.
	IdentifierFirst *bool `json:"identifier_first,omitempty"`
}

type PromptManager struct {
	*Management
}

type promptCustomText struct {
	Text *string
}

func newPromptManager(m *Management) *PromptManager {
	return &PromptManager{m}
}

// Read retrieves prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_prompts
func (m *PromptManager) Read(opts ...RequestOption) (p *Prompt, err error) {
	err = m.Request("GET", m.URI("prompts"), &p, opts...)
	return
}

// Update prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/patch_prompts
func (m *PromptManager) Update(p *Prompt, opts ...RequestOption) error {
	return m.Request("PATCH", m.URI("prompts"), p, opts...)
}

// MarshalJSON is a custom serializer for the promptCustomText type.
//
// We have to use one to avoid creating a new Request method that does not JSON-encode the request body.
func (t *promptCustomText) MarshalJSON() ([]byte, error) {
	return []byte(auth0.StringValue(t.Text)), nil
}

// UnmarshalJSON is a custom deserializer for the promptCustomText type.
//
// We have to use one to avoid creating a new Request method that does not JSON-decode the response body.
func (t *promptCustomText) UnmarshalJSON(b []byte) error {
	t.Text = auth0.String(string(b))
	return nil
}

// Retrieve custom text for a specific prompt and language.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) CustomText(p string, l string, opts ...RequestOption) (t *string, err error) {
	var r *promptCustomText
	err = m.Request("GET", m.URI("prompts", p, "custom-text", l), &r, opts...)
	t = r.Text
	return
}

// Set custom text for a specific prompt. Existing texts will be overwritten.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) SetCustomText(p string, l string, b *string, opts ...RequestOption) (err error) {
	r := &promptCustomText{Text: b}
	err = m.Request("PUT", m.URI("prompts", p, "custom-text", l), r, opts...)
	//lint:ignore SA4006 because the only purpose of b is to  be assigned to
	b = r.Text
	return
}
