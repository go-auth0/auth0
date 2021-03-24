package management

import "time"

type Enrollment struct {
	// ID for this enrollment
	ID *string `json:"id,omitempty"`
	// Status of this enrollment. Can be 'pending' or 'confirmed'
	Status *string `json:"status,omitempty"`
	// Device name (only for push notification).
	Name *string `json:"name,omitempty"`
	// Device identifier. This is usually the phone identifier.
	Identifier *string `json:"identifier,omitempty"`
	// Phone number.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Enrollment date and time.
	EnrolledAt *time.Time `json:"enrolled_at,omitempty"`
	// Last authentication date and time.
	LastAuth *time.Time `json:"last_auth,omitempty"`
}

type MultiFactor struct {
	// States if this factor is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Guardian Factor name
	Name *string `json:"name,omitempty"`

	// For factors with trial limits (e.g. SMS) states if those limits have been exceeded
	TrialExpired *bool `json:"trial_expired,omitempty"`
}

type MultiFactorPolicies []string

type MultiFactorProvider struct {
	//One of auth0|twilio|phone-message-hook
	Provider *string `json:"provider,omitempty"`
}

type PhoneMessageTypes struct {
	MessageTypes *[]string `json:"message_types,omitempty"`
}
type MultiFactorSMSTemplate struct {
	// Message sent to the user when they are invited to enroll with a phone number
	EnrollmentMessage *string `json:"enrollment_message,omitempty"`

	// Message sent to the user when they are prompted to verify their account
	VerificationMessage *string `json:"verification_message,omitempty"`
}

type MultiFactorProviderAmazonSNS struct {
	// AWS Access Key ID
	AccessKeyID *string `json:"aws_access_key_id,omitempty"`

	// AWS Secret Access Key ID
	SecretAccessKeyID *string `json:"aws_secret_access_key,omitempty"`

	// AWS Region
	Region *string `json:"aws_region,omitempty"`

	// SNS APNS Platform Application ARN
	APNSPlatformApplicationARN *string `json:"sns_apns_platform_application_arn,omitempty"`

	// SNS GCM Platform Application ARN
	GCMPlatformApplicationARN *string `json:"sns_gcm_platform_application_arn,omitempty"`
}

type MultiFactorProviderTwilio struct {
	// From number
	From *string `json:"from,omitempty"`

	// Copilot SID
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`

	// Twilio Authentication token
	AuthToken *string `json:"auth_token,omitempty"`

	// Twilio SID
	SID *string `json:"sid,omitempty"`
}

type GuardianManager struct {
	Enrollment  *EnrollmentManager
	MultiFactor *MultiFactorManager
}

func newGuardianManager(m *Management) *GuardianManager {
	return &GuardianManager{
		&EnrollmentManager{m},
		&MultiFactorManager{m,
			&MultiFactorPhone{m},
			&MultiFactorSMS{m},
			&MultiFactorPush{m},
			&MultiFactorEmail{m},
			&MultiFactorDUO{m},
			&MultiFactorOTP{m},
		},
	}
}

type EnrollmentManager struct {
	*Management
}

// Get retrieves an enrollment (including its status and type).
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/get_enrollments_by_id
func (m *EnrollmentManager) Get(id string, opts ...RequestOption) (en *Enrollment, err error) {
	err = m.Request("GET", m.URI("guardian", "enrollments", id), &en, opts...)
	return
}

// Delete an enrollment to allow the user to enroll with multi-factor authentication again.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/delete_enrollments_by_id
func (m *EnrollmentManager) Delete(id string, opts ...RequestOption) (err error) {
	err = m.Request("DELETE", m.URI("guardian", "enrollments", id), nil, opts...)
	return
}

type MultiFactorManager struct {
	*Management
	Phone *MultiFactorPhone
	SMS   *MultiFactorSMS
	Push  *MultiFactorPush
	Email *MultiFactorEmail
	DUO   *MultiFactorDUO
	OTP   *MultiFactorOTP
}

// Retrieves all factors.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/get_factors
func (m *MultiFactorManager) List(opts ...RequestOption) (mf []*MultiFactor, err error) {
	err = m.Request("GET", m.URI("guardian", "factors"), &mf, opts...)
	return
}

// Get MFA policies
//
// See: https://auth0.com/docs/api/management/v2/#!/Guardian/get_policies
func (m *MultiFactorManager) Policy(opts ...RequestOption) (p *MultiFactorPolicies, err error) {
	err = m.Request("GET", m.URI("guardian", "policies"), p, opts...)
	return
}

// Update MFA policies
//
// See: https://auth0.com/docs/api/management/v2/#!/Guardian/put_policies
//Expects an array of either ["all-applications"] or ["confidence-score"]
func (m *MultiFactorManager) UpdatePolicy(p *MultiFactorPolicies, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "policies"), p, opts...)
}

type MultiFactorPhone struct{ *Management }

// Update MFA Phone to be enabled.
// See: https://auth0.com/docs/api/management/v2/#!/Guardian/put_factors_by_name
func (m *MultiFactorPhone) Enable(enabled bool, opts ...RequestOption) error {
	// An endpoint for enabling Phone doesn't exist yet so we go towards
	// sms endpoint to be consistent with the other methods available for this struct.
	return m.Request("PUT", m.URI("guardian", "factors", "sms"), &MultiFactor{
		Enabled: &enabled,
	}, opts...)
}

// Retrieves the MFA Phone provider, one of ["auth0" or "twilio" or "phone-message-hook"]
// See: https://auth0.com/docs/api/management/v2/#!/Guardian/get_selected_provider
func (m *MultiFactorPhone) Provider(opts ...RequestOption) (p *MultiFactorProvider, err error) {
	err = m.Request("GET", m.URI("guardian", "factors", "phone", "selected-provider"), &p, opts...)
	return
}

// Update MFA Phone provider, one of ["auth0" or "twilio" or "phone-message-hook"]
// See: https://auth0.com/docs/api/management/v2/#!/Guardian/put_selected_provider
func (m *MultiFactorPhone) UpdateProvider(p *MultiFactorProvider, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "phone", "selected-provider"), &p, opts...)
}

// Retrieves the MFA Phone Message Type
// See: https://auth0.com/docs/api/management/v2/#!/Guardian/get_message_types
func (m *MultiFactorPhone) MessageTypes(opts ...RequestOption) (mt *PhoneMessageTypes, err error) {
	err = m.Request("GET", m.URI("guardian", "factors", "phone", "message-types"), &mt, opts...)
	return
}

// Update MFA Phone Message Type
// See: https://auth0.com/docs/api/management/v2/#!/Guardian/put_message_types
func (m *MultiFactorPhone) UpdateMessageTypes(mt *PhoneMessageTypes, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "phone", "message-types"), &mt, opts...)
}

type MultiFactorSMS struct{ *Management }

// Enable enables or disables the SMS Multi-factor Authentication.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_factors_by_name
func (m *MultiFactorSMS) Enable(enabled bool, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "sms"), &MultiFactor{
		Enabled: &enabled,
	}, opts...)
}

// Template retrieves enrollment and verification templates. You can use this to
// check the current values for your templates.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/get_templates
func (m *MultiFactorSMS) Template(opts ...RequestOption) (t *MultiFactorSMSTemplate, err error) {
	err = m.Request("GET", m.URI("guardian", "factors", "sms", "templates"), &t, opts...)
	return
}

// UpdateTemplate updates the enrollment and verification templates. It's useful
// to send custom messages on SMS enrollment and verification.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_templates
func (m *MultiFactorSMS) UpdateTemplate(t *MultiFactorSMSTemplate, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "sms", "templates"), t, opts...)
}

// Twilio returns the Twilio provider configuration.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/get_twilio
func (m *MultiFactorSMS) Twilio(opts ...RequestOption) (t *MultiFactorProviderTwilio, err error) {
	err = m.Request("GET", m.URI("guardian", "factors", "sms", "providers", "twilio"), &t, opts...)
	return
}

// UpdateTwilio updates the Twilio provider configuration.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_twilio
func (m *MultiFactorSMS) UpdateTwilio(t *MultiFactorProviderTwilio, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "sms", "providers", "twilio"), t, opts...)
}

type MultiFactorPush struct{ *Management }

// Enable enables or disables the Push Notification (via Auth0 Guardian)
// Multi-factor Authentication.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_factors_by_name
func (m *MultiFactorPush) Enable(enabled bool, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "push-notification"), &MultiFactor{
		Enabled: &enabled,
	}, opts...)
}

// AmazonSNS returns the Amazon Web Services (AWS) Simple Notification Service
// (SNS) provider configuration.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/get_sns
func (m *MultiFactorPush) AmazonSNS(opts ...RequestOption) (s *MultiFactorProviderAmazonSNS, err error) {
	err = m.Request("GET", m.URI("guardian", "factors", "push-notification", "providers", "sns"), &s, opts...)
	return
}

// UpdateAmazonSNS updates the Amazon Web Services (AWS) Simple Notification
// Service (SNS) provider configuration.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_sns
func (m *MultiFactorPush) UpdateAmazonSNS(sc *MultiFactorProviderAmazonSNS, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "push-notification", "providers", "sns"), sc, opts...)
}

type MultiFactorEmail struct{ *Management }

// Enable enables or disables the Email Multi-factor Authentication.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_factors_by_name
func (m *MultiFactorEmail) Enable(enabled bool, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "email"), &MultiFactor{
		Enabled: &enabled,
	}, opts...)
}

type MultiFactorDUO struct{ *Management }

// Enable enables or disables DUO Security Multi-factor Authentication.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_factors_by_name
func (m *MultiFactorDUO) Enable(enabled bool, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "duo"), &MultiFactor{
		Enabled: &enabled,
	}, opts...)
}

type MultiFactorOTP struct{ *Management }

// Enable enables or disables One-time Password Multi-factor Authentication.
//
// See: https://auth0.com/docs/api/management/v2#!/Guardian/put_factors_by_name
func (m *MultiFactorOTP) Enable(enabled bool, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("guardian", "factors", "otp"), &MultiFactor{
		Enabled: &enabled,
	}, opts...)
}
