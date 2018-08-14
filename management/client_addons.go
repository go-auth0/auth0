package management

type Samlp struct {
	Audience                       string        `json:"audience,omitempty"`
	Recipient                      string        `json:"recipient,omitempty"`
	Mappings                       *SamlMappings `json:"mappings,omitempty"`
	CreateUpnClaim                 bool          `json:"createUpnClaim" default:"true"`
	PassthroughClaimsWithNoMapping bool          `json:"passthroughClaimsWithNoMapping" default:"true"`
	MapUnknownClaimsAsIs           bool          `json:"mapUnknownClaimsAsIs"`
	MapIdentities                  bool          `json:"mapIdentities" default:"true"`
	SignatureAlgorithm             string        `json:"signatureAlgorithm,omitempty"`
	DigestAlgorithm                string        `json:"digestAlgorithm,omitempty"`
	Destination                    string        `json:"destination,omitempty"`
	LifetimeInSeconds              int           `json:"lifetimeInSeconds,omitempty"`
	SignResponse                   bool          `json:"signResponse"`
	TypedAttributes                bool          `json:"typedAttributes" default:"true"`
	IncludeAttributeNameFormat     bool          `json:"includeAttributeNameFormat" default:"true"`
	NameIdentifierFormat           string        `json:"nameIdentifierFormat,omitempty"`
	NameIdentifierProbes           []string      `json:"nameIdentifierProbes"`
	AuthnContextClassRef           string        `json:"authnContextClassRef,omitempty"`
	Logout                         *SamlLogout   `json:"logout,omitempty"`
	Binding                        string        `json:"binding,omitempty"`
	ApplicationCallbackUrl		   string        `json:"application_callback_url,omitempty"`
}

type SamlMappings struct {
	UserId     string `json:"user_id,omitempty"`
	Email      string `json:"email,omitempty"`
	Name       string `json:"name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	UPN        string `json:"upn,omitempty"`
	Groups     string `json:"groups,omitempty"`
}

type SamlLogout struct {
	Callback   string `json:"callback,omitempty"`
	SloEnabled bool   `json:"slo_enabled" default:"true"`
}
