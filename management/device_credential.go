package management

type DeviceCredential struct {

	// The id of the device credential
	ID *string `json:"id,omitempty"`

	// The device name.
	DeviceName *string `json:"device_name,omitempty"`

	// The device id.
	DeviceId *string `json:"device_id,omitempty"`

	// The device credential type. Can be either `public_key`,
	// `refresh_token` or `rotating_refresh_token`
	Type *string `json:"type,omitempty"`

	// The user id.
	UserId *string `json:"user_id,omitempty"`

	// The client id.
	ClientId *string `json:"client_id,omitempty"`
}

type DevicePublicKey struct {

	// The value of the device credential
	Value *string `json:"value,omitempty"`

	// The device name.
	DeviceName *string `json:"device_name,omitempty"`

	// The device id.
	DeviceId *string `json:"device_id,omitempty"`

	// The device credential type. Must be `public_key`.
	Type *string `json:"type,omitempty"`

	// The client id.
	ClientId *string `json:"client_id,omitempty"`
}

type DeviceCredentialManager struct {
	*Management
}

func newDeviceCredentialManager(m *Management) *DeviceCredentialManager {
	return &DeviceCredentialManager{m}
}

// Create a device public key credential.
//
// See: https://auth0.com/docs/api/management/v2#!/Device_Credentials/post_device_credentials
func (m *DeviceCredentialManager) Create(c *CustomDomain, opts ...RequestOption) (err error) {
	return m.Request("POST", m.URI("custom-domains"), c, opts...)
}

// Retrieve device credential details for a given user_id.
//
// See: https://auth0.com/docs/api/management/v2#!/Device_Credentials/get_device_credentials
func (m *CustomDomainManager) Get(userId string, opts ...RequestOption) (dc []*DeviceCredential, err error) {
	err = m.Request("GET", m.URI("device-credentials"), &dc, opts...)
	return
}

// Run the verification process on a custom domain.
//
// See: https://auth0.com/docs/api/management/v2#!/Device_Credentials/delete_device_credentials_by_id
func (m *DeviceCredentialManager) Delete(id string, opts ...RequestOption) (c *CustomDomain, err error) {
	err = m.Request("POST", m.URI("device-credentials", id), &c, opts...)
	return
}
