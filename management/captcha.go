package management

type CaptchaManager struct {
	*Management
}

// TODO: This is currently an un-documented API.

func newCaptchaManager(m *Management) *CaptchaManager {
	return &CaptchaManager{m}
}

type CaptchaSettings struct {
	Selected  string            `json:"selected"`
	Policy    string            `json:"policy"`
	Providers *CaptchaProviders `json:"providers"`
}

type CaptchaProviders struct {
	Auth0               CaptchaProviderAuth0               `json:"auth0"`
	RecaptchaV2         CaptchaProviderRecaptchaV2         `json:"recaptcha_v2"`
	RecaptchaEnterprise CaptchaProviderRecaptchaEnterprise `json:"recaptcha_enterprise"`
}

type CaptchaProviderAuth0 struct{}

type CaptchaProviderRecaptchaV2 struct {
	SiteKey string `json:"siteKey"`
	Secret  string `json:"secret"`
}

type CaptchaProviderRecaptchaEnterprise struct {
	SiteKey   string `json:"siteKey"`
	APIKey    string `json:"apiKey"`
	ProjectID string `json:"projectId"`
}

// Get captcha settings for the auth0 tenant
func (m *CaptchaManager) GetCaptchaSettings(opts ...RequestOption) (c *CaptchaSettings, err error) {
	err = m.Request("GET", m.URI("anomaly", "captchas"), &c, opts...)
	return
}

// Sets the captcha settings for the auth0 tenant
func (m *CaptchaManager) SetCaptchaSettings(c *CaptchaSettings, opts ...RequestOption) (err error) {
	return m.Request("POST", m.URI("anomaly", "captchas"), c, opts...)
}
