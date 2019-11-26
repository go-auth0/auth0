package management

type BlacklistToken struct {

	// The "aud" (audience) claim identifies the recipients that the JWT is
	// intended for.
	//
	// See: https://tools.ietf.org/html/rfc7519#section-4.1.3
	Audience string `json:"aud,omitempty"`

	// The "jti" (JWT ID) claim provides a unique (within "aud") identifier for
	// the JWT.
	//
	// See: https://tools.ietf.org/html/rfc7519#section-4.1.7
	JTI string `json:"jti,omitempty"`
}

type BlacklistManager struct {
	m *Management
}

func NewBlacklistManager(m *Management) *BlacklistManager {
	return &BlacklistManager{m}
}

func (bm *BlacklistManager) List() (bl []*BlacklistToken, err error) {
	err = bm.m.get(bm.m.uri("blacklists", "tokens"), &bl)
	return
}

func (bm *BlacklistManager) Update(bt *BlacklistToken) error {
	return bm.m.post(bm.m.uri("blacklists", "tokens"), bt)
}
