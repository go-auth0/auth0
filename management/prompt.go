package management

type Prompt struct {
	// Which login experience to use. Can be `new` or `classic`.
	UniversalLoginExperience string `json:"universal_login_experience,omitempty"`
}

type PromptManager struct {
	m *Management
}

func NewPromptManager(m *Management) *PromptManager {
	return &PromptManager{m}
}

// Read retrieves prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_prompts
func (pm *PromptManager) Read() (*Prompt, error) {
	p := new(Prompt)
	err := pm.m.get(pm.m.uri("prompts"), p)
	return p, err
}

// Update prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/patch_prompts
func (pm *PromptManager) Update(p *Prompt) error {
	return pm.m.patch(pm.m.uri("prompts"), p)
}
