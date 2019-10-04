package management

type PromptSettings struct {
	// Which login experience to use. Can be `new` or `classic`.
	UniversalLoginExperience string `json:"universal_login_experience,omitempty"`
}

type PromptManager struct {
	m *Management
}

func NewPromptManager(m *Management) *PromptManager {
	return &PromptManager{m}
}

func (pm *PromptManager) GetSettings() (*PromptSettings, error) {
	ps := new(PromptSettings)
	err := pm.m.get(pm.m.uri("prompts"), ps)
	return ps, err
}

func (pm *PromptManager) UpdateSettings(ps *PromptSettings) error {
	return pm.m.patch(pm.m.uri("prompts"), ps)
}
