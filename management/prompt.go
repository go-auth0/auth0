package management

import (
	"encoding/json"
)

const (
	PromptConsent = "consent"
	PromptDeviceFlow = "device-flow"
	PromptEmailOtpChallengeFlow = "email-otp-challenge"
	PromptEmailVerificationFlow = "email-verification"
	PromptLogin = "login"
	PromptLoginEmailVerification = "login-email-verification"
	PromptMfa = "mfa"
	PromptMfaEmail = "mfa-email"
	PromptMfaOtp = "mfa-otp"
	PromptMfaPhone = "mfa-phone"
	PromptMfaPush = "mfa-push"
	PromptMfaRecoveryCode = "mfa-recovery-code"
	PromptMfaSms = "mfa-sms"
	PromptMfaVoice = "mfa-voice"
	PromptResetPassword = "reset-password"
	PromptSignup = "signup"
)

type Prompt struct {
	// Which login experience to use. Can be `new` or `classic`.
	UniversalLoginExperience string `json:"universal_login_experience,omitempty"`
}

type PromptCustomText struct {
	Language string `json:"-"`
	Prompt string `json:"-"`
	Screens interface{} `json:"-"`
}

type PromptManager struct {
	*Management
}

type ConsentScreens struct {
	Consent map[string]interface{} `json:"consent,omitempty"`
}
type DeviceFlowScreens struct {
	DeviceCodeActivation map[string]interface{} `json:"device-code-activation,omitempty"`
	DeviceCodeActivationAllowed map[string]interface{} `json:"device-code-activation-allowed,omitempty"`
	DeviceCodeActivationDenied map[string]interface{} `json:"device-code-activation-denied,omitempty"`
	DeviceCodeActivationConfirmation map[string]interface{} `json:"device-code-activation-confirmation,omitempty"`
}
type EmailOtpChallengeScreens struct {
	EmailOtpChallenge map[string]interface{} `json:"email-otp-challenge,omitempty"`
}
type EmailVerificationScreens struct {
	EmailVerificationResult map[string]interface{} `json:"email-verification-result,omitempty"`
}
type LoginScreens struct {
	Login map[string]interface{} `json:"login,omitempty"`
}
type LoginEmailVerificationScreens struct {
	LoginEmailVerification map[string]interface{} `json:"login-email-verification,omitempty"`
}
type MfaScreens struct {
	MfaEnrollResult map[string]interface{} `json:"mfa-enroll-result,omitempty"`
	MfaLoginOptions map[string]interface{} `json:"mfa-login-options,omitempty"`
}
type MfaEmailScreens struct {
	MfaEmailChallenge map[string]interface{} `json:"mfa-email-challenge,omitempty"`
	MfaEmailList map[string]interface{} `json:"mfa-email-list,omitempty"`
}
type MfaOtpScreens struct {
	MfaOtpEnrollmentQr map[string]interface{} `json:"mfa-otp-enrollment-qr,omitempty"`
	MfaOtpEnrollmentCode map[string]interface{} `json:"mfa-otp-enrollment-code,omitempty"`
	MfaOtpChallenge map[string]interface{} `json:"mfa-otp-challenge,omitempty"`
}
type MfaPhoneScreens struct {
	MfaPhoneChallenge map[string]interface{} `json:"mfa-phone-challenge,omitempty"`
	MfaPhoneEnrollment map[string]interface{} `json:"mfa-phone-enrollment,omitempty"`
}
type MfaPushScreens struct {
	MfaPushWelcome map[string]interface{} `json:"mfa-push-welcome,omitempty"`
	MfaPushEnrollmentQr map[string]interface{} `json:"mfa-push-enrollment-qr,omitempty"`
	MfaPushChallengePush map[string]interface{} `json:"mfa-push-challenge-push,omitempty"`
	MfaPushChallengeCode map[string]interface{} `json:"mfa-push-challenge-code,omitempty"`
	MfaPushList map[string]interface{} `json:"mfa-push-list,omitempty"`
}
type MfaRecoveryCodeScreens struct {
	MfaRecoveryCodeEnrollment map[string]interface{} `json:"mfa-recovery-code-enrollment,omitempty"`
	MfaRecoveryCodeChallenge map[string]interface{} `json:"mfa-recovery-code-challenge,omitempty"`
}
type MfaSmsScreens struct {
	MfaCountryCodes map[string]interface{} `json:"mfa-country-codes,omitempty"`
	MfaSmsEnrollment map[string]interface{} `json:"mfa-sms-enrollment,omitempty"`
	MfaSmsChallenge map[string]interface{} `json:"mfa-sms-challenge,omitempty"`
	MfaSmsList map[string]interface{} `json:"mfa-sms-list,omitempty"`
}
type MfaVoiceScreens struct {
	MfaVoiceEnrollment map[string]interface{} `json:"mfa-voice-enrollment,omitempty"`
	MfaVoiceChallenge map[string]interface{} `json:"mfa-voice-challenge,omitempty"`
}
type ResetPasswordScreens struct {
	ResetPasswordRequest map[string]interface{} `json:"reset-password-request,omitempty"`
	ResetPasswordEmail map[string]interface{} `json:"reset-password-email,omitempty"`
	ResetPassword map[string]interface{} `json:"reset-password,omitempty"`
	ResetPasswordError map[string]interface{} `json:"reset-password-error,omitempty"`
}
type SignupScreens struct {
	Signup map[string]interface{} `json:"signup,omitempty"`
}

func newPromptManager(m *Management) *PromptManager {
	return &PromptManager{m}
}

func (pct *PromptCustomText) MarshalJSON() ([]byte, error) {
	return json.Marshal(pct.Screens)
}

func (pct *PromptCustomText) UnmarshalJSON(b []byte) error {
	var v interface{}

	switch pct.Prompt {
	case PromptConsent:
		v = &ConsentScreens{}
	case PromptDeviceFlow:
		v = &DeviceFlowScreens{}
	case PromptEmailOtpChallengeFlow:
		v = &EmailOtpChallengeScreens{}
	case PromptEmailVerificationFlow:
		v = &EmailVerificationScreens{}
	case PromptLogin:
		v = &LoginScreens{}
	case PromptLoginEmailVerification:
		v = &LoginEmailVerificationScreens{}
	case PromptMfa:
		v = &MfaScreens{}
	case PromptMfaEmail:
		v = &MfaEmailScreens{}
	case PromptMfaOtp:
		v = &MfaOtpScreens{}
	case PromptMfaPhone:
		v = &MfaPhoneScreens{}
	case PromptMfaPush:
		v = &MfaPushScreens{}
	case PromptMfaRecoveryCode:
		v = &MfaRecoveryCodeScreens{}
	case PromptMfaSms:
		v = &MfaSmsScreens{}
	case PromptMfaVoice:
		v = &MfaVoiceScreens{}
	case PromptResetPassword:
		v = &ResetPasswordScreens{}
	case PromptSignup:
		v = &SignupScreens{}
	default:
		v = make(map[string]interface{})
	}

	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	pct.Screens = v

	return nil
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

// ReadCustomText retrieve custom text for a specific prompt and language.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadCustomText(prompt string, language string, opts ...RequestOption) (pct *PromptCustomText, err error) {
	pct = &PromptCustomText{
		Prompt:   prompt,
		Language: language,
	}
	err = m.Request("GET", m.URI("prompts", prompt, "custom-text", language), &pct, opts...)
	return
}

// UpdateCustomText set custom text for a specific prompt. Existing texts will be overwritten.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) UpdateCustomText(pct *PromptCustomText, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", pct.Prompt, "custom-text", pct.Language), pct, opts...)
}
