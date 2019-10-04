package management

type BlacklistToken struct {

	// JWT's aud claim (the client_id to which the JWT was issued).
	Aud string `json:"aud,omitempty"`

	// jti (unique ID within aud) of the blacklisted JWT.
	Jti string `json:"jti,omitempty"`
}

type BlacklistManager struct {
	m *Management
}

func NewBlacklistManager(m *Management) *BlacklistManager {
	return &BlacklistManager{m}
}

func (bm *BlacklistManager) GetBlacklistedTokens() (bl []*BlacklistToken, err error) {
	err = bm.m.get(bm.m.uri("blacklists/tokens"), &bl)
	return
}

func (bm *BlacklistManager) BlacklistToken(bt *BlacklistToken) error {
	return bm.m.post(bm.m.uri("blacklists/tokens"), bt)
}
