package management

import "encoding/json"

type MultiFactor struct {
	// States if this factor is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Guardian Factor name
	Name *string `json:"name,omitempty"`

	// For factors with trial limits (e.g. SMS) states if those limits have been exceeded
	TrialExpired *bool `json:"trial_expired,omitempty"`
}

func (mfa *MultiFactor) String() string {
	b, _ := json.MarshalIndent(mfa, "", "  ")
	return string(b)
}

type MultiFactorSMSTemplate struct {
	// Message sent to the user when they are invited to enroll with a phone number
	EnrollmentMessage *string `json:"enrollment_message,omitempty"`

	// Message sent to the user when they are prompted to verify their account
	VerificationMessage *string `json:"verification_message,omitempty"`
}

func (sms *MultiFactorSMSTemplate) String() string {
	b, _ := json.MarshalIndent(sms, "", "  ")
	return string(b)
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

func (sns *MultiFactorProviderAmazonSNS) String() string {
	b, _ := json.MarshalIndent(sns, "", "  ")
	return string(b)
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

func (twilio *MultiFactorProviderTwilio) String() string {
	b, _ := json.MarshalIndent(twilio, "", "  ")
	return string(b)
}

type GuardianManager struct {
	MultiFactor *MultiFactorManager
}

func NewGuardianManager(m *Management) *GuardianManager {
	return &GuardianManager{
		&MultiFactorManager{m,
			&MultiFactorSMS{m},
			&MultiFactorPush{m},
			&MultiFactorEmail{m},
			&MultiFactorDUO{m},
			&MultiFactorOTP{m},
		},
	}
}

type MultiFactorManager struct {
	m *Management

	SMS   *MultiFactorSMS
	Push  *MultiFactorPush
	Email *MultiFactorEmail
	DUO   *MultiFactorDUO
	OTP   *MultiFactorOTP
}

func (mfm *MultiFactorManager) List(opts ...reqOption) ([]*MultiFactor, error) {
	var mf []*MultiFactor
	err := mfm.m.get(mfm.m.uri("guardian", "factors")+mfm.m.q(opts), &mf)
	return mf, err
}

type MultiFactorSMS struct{ m *Management }

func (sm *MultiFactorSMS) Enable(enabled bool) error {
	return sm.m.put(sm.m.uri("guardian", "factors", "sms"), &MultiFactor{
		Enabled: &enabled,
	})
}

func (sm *MultiFactorSMS) Template() (*MultiFactorSMSTemplate, error) {
	t := new(MultiFactorSMSTemplate)
	err := sm.m.get(sm.m.uri("guardian", "factors", "sms", "templates"), t)
	return t, err
}

func (sm *MultiFactorSMS) UpdateTemplate(st *MultiFactorSMSTemplate) error {
	return sm.m.put(sm.m.uri("guardian", "factors", "sms", "templates"), st)
}

func (sm *MultiFactorSMS) Twilio() (*MultiFactorProviderTwilio, error) {
	tc := new(MultiFactorProviderTwilio)
	err := sm.m.get(sm.m.uri("guardian", "factors", "sms", "providers", "twilio"), tc)
	return tc, err
}

func (sm *MultiFactorSMS) UpdateTwilio(twilio *MultiFactorProviderTwilio) error {
	return sm.m.put(sm.m.uri("guardian", "factors", "sms", "providers", "twilio"), twilio)
}

type MultiFactorPush struct{ m *Management }

func (pm *MultiFactorPush) Enable(enabled bool) error {
	return pm.m.put(pm.m.uri("guardian", "factors", "push-notification"), &MultiFactor{
		Enabled: &enabled,
	})
}

func (pm *MultiFactorPush) AmazonSNS() (*MultiFactorProviderAmazonSNS, error) {
	sc := new(MultiFactorProviderAmazonSNS)
	err := pm.m.get(pm.m.uri("guardian", "factors", "push-notification", "providers", "sns"), sc)
	return sc, err
}

func (pm *MultiFactorPush) UpdateAmazonSNS(sc *MultiFactorProviderAmazonSNS) error {
	return pm.m.put(pm.m.uri("guardian", "factors", "push-notification", "providers", "sns"), sc)
}

type MultiFactorEmail struct{ m *Management }

func (em *MultiFactorEmail) Enable(enabled bool) error {
	return em.m.put(em.m.uri("guardian", "factors", "email"), &MultiFactor{
		Enabled: &enabled,
	})
}

type MultiFactorDUO struct{ m *Management }

func (em *MultiFactorDUO) Enable(enabled bool) error {
	return em.m.put(em.m.uri("guardian", "factors", "duo"), &MultiFactor{
		Enabled: &enabled,
	})
}

type MultiFactorOTP struct{ m *Management }

func (em *MultiFactorOTP) Enable(enabled bool) error {
	return em.m.put(em.m.uri("guardian", "factors", "otp"), &MultiFactor{
		Enabled: &enabled,
	})
}
