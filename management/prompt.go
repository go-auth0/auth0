package management

////////////////////////////////////////////////
// Get prompts settings, Update prompts settings
//

type Prompt struct {
	// Which login experience to use. Can be `new` or `classic`.
	UniversalLoginExperience string `json:"universal_login_experience,omitempty"`

	// IdentifierFirst determines if the login screen prompts for just the identifier, identifier and password first.
	IdentifierFirst *bool `json:"identifier_first,omitempty"`
}

type PromptManager struct {
	*Management
}

func newPromptManager(m *Management) *PromptManager {
	return &PromptManager{m}
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

// PromptConsent stores consent custom text
//
// See: https://auth0.com/docs/universal-login/prompt-consent

type ScreenConsent struct {
	PageTitle              string `json:"pageTitle,omitempty"`
	Title                  string `json:"title,omitempty"`
	PickerTitle            string `json:"pickerTitle,omitempty"`
	MessageMultipleTenants string `json:"messageMultipleTenants,omitempty"`
	AudiencePickerAltText  string `json:"audiencePickerAltText,omitempty"`
	MessageSingleTenant    string `json:"messageSingleTenant,omitempty"`
	AcceptButtonText       string `json:"acceptButtonText,omitempty"`
	DeclineButtonText      string `json:"declineButtonText,omitempty"`
	InvalidAction          string `json:"invalid-action,omitempty"`
	InvalidAudience        string `json:"invalid-audience,omitempty"`
	InvalidScope           string `json:"invalid-scope,omitempty"`
}

type PromptConsent struct {
	Language string `json:"language,omitempty"`
	Consent *ScreenConsent `json:"consent,omitempty"`
}

// ReadPromptConsent retrieves consent custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptConsent(language string, opts ...RequestOption) (p *PromptConsent, err error) {
	err = m.Request("GET", m.URI("prompts", "consent", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptConsent replaces consent custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptConsent(language string, p *PromptConsent, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "consent", "custom-text", language), p, opts...)
}

// PromptDeviceFlow stores device-flow custom text
//
// See: https://auth0.com/docs/universal-login/prompt-device-flow

type ScreenDeviceCodeActivation struct {
	PageTitle          string `json:"pageTitle,omitempty"`
	ButtonText         string `json:"buttonText,omitempty"`
	Description        string `json:"description,omitempty"`
	Placeholder        string `json:"placeholder,omitempty"`
	Title              string `json:"title,omitempty"`
	InvalidExpiredCode string `json:"invalid-expired-code,omitempty"`
	NoCode             string `json:"no-code,omitempty"`
	InvalidCode        string `json:"invalid-code,omitempty"`
}

type ScreenDeviceCodeActivationAllowed struct {
	PageTitle   string `json:"pageTitle,omitempty"`
	Description string `json:"description,omitempty"`
	EventTitle  string `json:"eventTitle,omitempty"`
}

type ScreenDeviceCodeActivationDenied struct {
	PageTitle   string `json:"pageTitle,omitempty"`
	Description string `json:"description,omitempty"`
	EventTitle  string `json:"eventTitle,omitempty"`
}

type ScreenDeviceCodeConfirmation struct {
	PageTitle         string `json:"pageTitle,omitempty"`
	Description       string `json:"description,omitempty"`
	InputCodeLabel    string `json:"inputCodeLabel,omitempty"`
	Title             string `json:"title,omitempty"`
	ConfirmButtonText string `json:"confirmButtonText,omitempty"`
	CancelButtonText  string `json:"cancelButtonText,omitempty"`
	ConfirmationText  string `json:"confirmationText,omitempty"`
}

type PromptDeviceFlow struct {
	DeviceCodeActivation        *ScreenDeviceCodeActivation        `json:"device-code-activation,omitempty"`
	DeviceCodeActivationAllowed *ScreenDeviceCodeActivationAllowed `json:"device-code-activation-allowed,omitempty"`
	DeviceCodeActivationDenied  *ScreenDeviceCodeActivationDenied  `json:"device-code-activation-denied,omitempty"`
	DeviceCodeConfirmation      *ScreenDeviceCodeConfirmation      `json:"device-code-confirmation"`
}

// ReadPromptDeviceFlow retrieves device-flow custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptDeviceFlow(language string, opts ...RequestOption) (p *PromptDeviceFlow, err error) {
	err = m.Request("GET", m.URI("prompts", "device-flow", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptDeviceFlow replaces device-flow custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptDeviceFlow(language string, p *PromptDeviceFlow, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "device-flow", "custom-text", language), p, opts...)
}

// PromptEmailOtpChallenge stores email-otp-challenge custom text
//
// See: https://auth0.com/docs/universal-login/prompt-email-otp-challenge

type ScreenEmailOtpChallenge struct {
	PageTitle            string `json:"pageTitle,omitempty"`
	ButtonText           string `json:"buttonText,omitempty"`
	Description          string `json:"description,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
	ResendActionText     string `json:"resendActionText,omitempty"`
	ResendText           string `json:"resendText,omitempty"`
	Title                string `json:"title,omitempty"`
	InvalidOtpCodeFormat string `json:"invalid-otp-code-format,omitempty"`
	InvalidCode          string `json:"invalid-code,omitempty"`
	InvalidExpiredCode   string `json:"invalid-expired-code,omitempty"`
	AuthenticatorError   string `json:"authenticator-error,omitempty"`
	TooManyEmail         string `json:"too-many-email,omitempty"`
}

type PromptEmailOtpChallenge struct {
	EmailOtpChallenge *ScreenEmailOtpChallenge `json:"email-otp-challenge,omitempty"`
}

// ReadPromptEmailOtpChallenge retrieves email-otp-challenge custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptEmailOtpChallenge(language string, opts ...RequestOption) (p *PromptEmailOtpChallenge, err error) {
	err = m.Request("GET", m.URI("prompts", "email-otp-challenge", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptEmailOtpChallenge replaces email-otp-challenge custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptEmailOtpChallenge(language string, p *PromptEmailOtpChallenge, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "email-otp-challenge", "custom-text", language), p, opts...)
}

// PromptEmailVerification stores email-verification custom text
//
// See: https://auth0.com/docs/universal-login/prompt-email-verification

type ScreenEmailVerificationResult struct {
	PageTitle                       string `json:"pageTitle,omitempty"`
	VerifiedTitle                   string `json:"verifiedTitle,omitempty"`
	ErrorTitle                      string `json:"errorTitle,omitempty"`
	VerifiedDescription             string `json:"verifiedDescription,omitempty"`
	AlreadyVerifiedDescription      string `json:"alreadyVerifiedDescription,omitempty"`
	InvalidAccountOrCodeDescription string `json:"invalidAccountOrCodeDescription,omitempty"`
	UnknownErrorDescription         string `json:"unknownErrorDescription,omitempty"`
	ButtonText                      string `json:"buttonText,omitempty"`
	Auth0UsersExpiredTicket         string `json:"auth0-users-expired-ticket,omitempty"`
	CustomScriptErrorCode           string `json:"custom-script-error-code,omitempty"`
	Auth0UsersUsedTicket            string `json:"auth0-users-used-ticket,omitempty"`
	Auth0UsersValidation            string `json:"auth0-users-validation,omitempty"`
}

type PromptEmailVerification struct {
	EmailVerificationResult *ScreenEmailVerificationResult `json:"email-verification-result,omitempty"`
}

// ReadPromptEmailVerification retrieves email-verification custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptEmailVerification(language string, opts ...RequestOption) (p *PromptEmailVerification, err error) {
	err = m.Request("GET", m.URI("prompts", "email-verification", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptEmailVerification replaces email-verification custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptEmailVerification(language string, p *PromptEmailVerification, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "email-verification", "custom-text", language), p, opts...)
}

// PromptAcceptInvitation stores accept-invitation custom text
//
// See: https://auth0.com/docs/universal-login/prompt-accept-invitation

type ScreenAcceptInvitation struct {
	PageTitle   string `json:"pageTitle,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ButtonText  string `json:"buttonText,omitempty"`
	LogoAltText string `json:"logoAltText,omitempty"`
}

type PromptAcceptInvitation struct {
	AcceptInvitation *ScreenAcceptInvitation `json:"accept-invitation,omitempty"`
}

// ReadPromptAcceptInvitation retrieves invitation custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptAcceptInvitation(language string, opts ...RequestOption) (p *PromptAcceptInvitation, err error) {
	err = m.Request("GET", m.URI("prompts", "invitation", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptAcceptInvitation replaces invitation custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptAcceptInvitation(language string, p *PromptAcceptInvitation, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "invitation", "custom-text", language), p, opts...)
}

// PromptLogin stores login custom text
//
// See: https://auth0.com/docs/universal-login/prompt-login

type ScreenLogin struct {
	PageTitle                     string `json:"pageTitle,omitempty"`
	Title                         string `json:"title,omitempty"`
	Description                   string `json:"description,omitempty"`
	SeparatorText                 string `json:"separatorText,omitempty"`
	ButtonText                    string `json:"buttonText,omitempty"`
	FederatedConnectionButtonText string `json:"federatedConnectionButtonText,omitempty"`
	SignupActionLinkText          string `json:"signupActionLinkText,omitempty"`
	SignupActionText              string `json:"signupActionText,omitempty"`
	ForgotPasswordText            string `json:"forgotPasswordText,omitempty"`
	PasswordPlaceholder           string `json:"passwordPlaceholder,omitempty"`
	UsernamePlaceholder           string `json:"usernamePlaceholder,omitempty"`
	//Documented but not supported
	//CaptchaCodePlaceholder        string `json:"captchaCodePlaceholder,omitempty"`
	//CaptchaMatchExprPlaceholder   string `json:"captchaMatchExprPlaceholder,omitempty"`
	EmailPlaceholder      string `json:"emailPlaceholder,omitempty"`
	EditEmailText         string `json:"editEmailText,omitempty"`
	AlertListTitle        string `json:"alertListTitle,omitempty"`
	InvitationTitle       string `json:"invitationTitle,omitempty"`
	InvitationDescription string `json:"invitationDescription,omitempty"`
	WrongCredentials      string `json:"wrong-credentials,omitempty"`
	//Documented but not supported
	//WrongCaptcha                  string `json:"wrong-captcha,omitempty"`
	InvalidCode           string `json:"invalid-code,omitempty"`
	InvalidExpiredCode    string `json:"invalid-expired-code,omitempty"`
	InvalidEmailFormat    string `json:"invalid-email-format,omitempty"`
	WrongEmailCredentials string `json:"wrong-email-credentials,omitempty"`
	CustomScriptErrorCode string `json:"custom-script-error-code,omitempty"`
	Auth0UsersValidation  string `json:"auth0-users-validation,omitempty"`
	AuthenticationFailure string `json:"authentication-failure,omitempty"`
	InvalidConnection     string `json:"invalid-connection,omitempty"`
	IpBlocked             string `json:"ip-blocked,omitempty"`
	NoDbConnection        string `json:"no-db-connection,omitempty"`
	PasswordBreached      string `json:"password-breached,omitempty"`
	UserBlocked           string `json:"user-blocked,omitempty"`
	SameUserLogin         string `json:"same-user-login,omitempty"`
	NoEmail               string `json:"no-email,omitempty"`
	NoPassword            string `json:"no-password,omitempty"`
	NoUsername            string `json:"no-username,omitempty"`
}

type PromptLogin struct {
	Login *ScreenLogin `json:"login,omitempty"`
}

// ReadPromptLogin retrieves login custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptLogin(language string, opts ...RequestOption) (p *PromptLogin, err error) {
	err = m.Request("GET", m.URI("prompts", "login", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptLogin replaces login custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptLogin(language string, p *PromptLogin, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "login", "custom-text", language), p, opts...)
}

// PromptLoginId stores login-id custom text
//
// See: https://auth0.com/docs/universal-login/prompt-login-id

type ScreenLoginId struct {
	PageTitle                     string `json:"pageTitle,omitempty"`
	Title                         string `json:"title,omitempty"`
	Description                   string `json:"description,omitempty"`
	SeparatorText                 string `json:"separatorText,omitempty"`
	ButtonText                    string `json:"buttonText,omitempty"`
	FederatedConnectionButtonText string `json:"federatedConnectionButtonText,omitempty"`
	SignupActionLinkText          string `json:"signupActionLinkText,omitempty"`
	SignupActionText              string `json:"signupActionText,omitempty"`
	PasswordPlaceholder           string `json:"passwordPlaceholder,omitempty"`
	UsernamePlaceholder           string `json:"usernamePlaceholder,omitempty"`
	EmailPlaceholder              string `json:"emailPlaceholder,omitempty"`
	EditEmailText                 string `json:"editEmailText,omitempty"`
	AlertListTitle                string `json:"alertListTitle,omitempty"`
	LogoAltText                   string `json:"logoAltText,omitempty"`
	WrongCredentials              string `json:"wrong-credentials,omitempty"`
	InvalidCode                   string `json:"invalid-code,omitempty"`
	InvalidExpiredCode            string `json:"invalid-expired-code,omitempty"`
	InvalidEmailFormat            string `json:"invalid-email-format,omitempty"`
	WrongEmailCredentials         string `json:"wrong-email-credentials,omitempty"`
	CustomScriptErrorCode         string `json:"custom-script-error-code,omitempty"`
	Auth0UsersValidation          string `json:"auth0-users-validation,omitempty"`
	AuthenticationFailure         string `json:"authentication-failure,omitempty"`
	InvalidConnection             string `json:"invalid-connection,omitempty"`
	IpBlocked                     string `json:"ip-blocked,omitempty"`
	NoDbConnection                string `json:"no-db-connection,omitempty"`
	NoHrdConnection               string `json:"no-hrd-connection,omitempty"`
	PasswordBreached              string `json:"password-breached,omitempty"`
	UserBlocked                   string `json:"user-blocked,omitempty"`
	SameUserLogin                 string `json:"same-user-login,omitempty"`
	NoEmail                       string `json:"no-email,omitempty"`
	NoPassword                    string `json:"no-password,omitempty"`
	NoUsername                    string `json:"no-username,omitempty"`
}

type PromptLoginId struct {
	LoginId *ScreenLoginId `json:"login-id,omitempty"`
}

// ReadPromptLoginId retrieves login-id custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptLoginId(language string, opts ...RequestOption) (p *PromptLoginId, err error) {
	err = m.Request("GET", m.URI("prompts", "login-id", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptLoginId replaces login-id custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptLoginId(language string, p *PromptLoginId, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "login-id", "custom-text", language), p, opts...)
}

// PromptLoginPassword stores login-password custom text
//
// See: https://auth0.com/docs/universal-login/prompt-login-password

type ScreenLoginPassword struct {
	PageTitle                     string `json:"pageTitle,omitempty"`
	Title                         string `json:"title,omitempty"`
	Description                   string `json:"description,omitempty"`
	SeparatorText                 string `json:"separatorText,omitempty"`
	ButtonText                    string `json:"buttonText,omitempty"`
	FederatedConnectionButtonText string `json:"federatedConnectionButtonText,omitempty"`
	SignupActionLinkText          string `json:"signupActionLinkText,omitempty"`
	SignupActionText              string `json:"signupActionText,omitempty"`
	ForgotPasswordText            string `json:"forgotPasswordText,omitempty"`
	PasswordPlaceholder           string `json:"passwordPlaceholder,omitempty"`
	UsernamePlaceholder           string `json:"usernamePlaceholder,omitempty"`
	EmailPlaceholder              string `json:"emailPlaceholder,omitempty"`
	EditEmailText                 string `json:"editEmailText,omitempty"`
	AlertListTitle                string `json:"alertListTitle,omitempty"`
	InvitationTitle               string `json:"invitationTitle,omitempty"`
	InvitationDescription         string `json:"invitationDescription,omitempty"`
	LogoAltText                   string `json:"logoAltText,omitempty"`
	WrongCredentials              string `json:"wrong-credentials,omitempty"`
	InvalidCode                   string `json:"invalid-code,omitempty"`
	InvalidExpiredCode            string `json:"invalid-expired-code,omitempty"`
	InvalidEmailFormat            string `json:"invalid-email-format,omitempty"`
	WrongEmailCredentials         string `json:"wrong-email-credentials,omitempty"`
	CustomScriptErrorCode         string `json:"custom-script-error-code,omitempty"`
	Auth0UsersValidation          string `json:"auth0-users-validation,omitempty"`
	AuthenticationFailure         string `json:"authentication-failure,omitempty"`
	InvalidConnection             string `json:"invalid-connection,omitempty"`
	IpBlocked                     string `json:"ip-blocked,omitempty"`
	NoDbConnection                string `json:"no-db-connection,omitempty"`
	PasswordBreached              string `json:"password-breached,omitempty"`
	UserBlocked                   string `json:"user-blocked,omitempty"`
	SameUserLogin                 string `json:"same-user-login,omitempty"`
	NoEmail                       string `json:"no-email,omitempty"`
	NoPassword                    string `json:"no-password,omitempty"`
	NoUsername                    string `json:"no-username,omitempty"`
}

type PromptLoginPassword struct {
	LoginPassword *ScreenLoginPassword `json:"login-password,omitempty"`
}

// ReadPromptLoginPassword retrieves login-password custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptLoginPassword(language string, opts ...RequestOption) (p *PromptLoginPassword, err error) {
	err = m.Request("GET", m.URI("prompts", "login-password", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptLoginPassword replaces login-password custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptLoginPassword(language string, p *PromptLoginPassword, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "login-password", "custom-text", language), p, opts...)
}

// PromptLoginEmailVerification stores login-email-verification custom text
//
// See: https://auth0.com/docs/universal-login/prompt-login-email-verification

type ScreenLoginEmailVerification struct {
	PageTitle            string `json:"pageTitle,omitempty"`
	ButtonText           string `json:"buttonText,omitempty"`
	Description          string `json:"description,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
	ResendActionText     string `json:"resendActionText,omitempty"`
	ResendText           string `json:"resendText,omitempty"`
	Title                string `json:"title,omitempty"`
	InvalidOtpCodeFormat string `json:"invalid-otp-code-format,omitempty"`
	InvalidCode          string `json:"invalid-code,omitempty"`
	InvalidExpiredCode   string `json:"invalid-expired-code,omitempty"`
	AuthenticatorError   string `json:"authenticator-error,omitempty"`
	TooManyEmail         string `json:"too-many-email,omitempty"`
}

type PromptLoginEmailVerification struct {
	LoginEmailVerification *ScreenLoginEmailVerification `json:"login-email-verification,omitempty"`
}

// ReadPromptLoginEmailVerification retrieves login-email-verification custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptLoginEmailVerification(language string, opts ...RequestOption) (p *PromptLoginEmailVerification, err error) {
	err = m.Request("GET", m.URI("prompts", "login-email-verification", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptLoginEmailVerification replaces login-email-verification custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptLoginEmailVerification(language string, p *PromptLoginEmailVerification, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "login-email-verification", "custom-text", language), p, opts...)
}

// PromptMfa stores mfa custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa

type ScreenMfaEnrollResult struct {
	PageTitle                  string `json:"pageTitle,omitempty"`
	EnrolledTitle              string `json:"enrolledTitle,omitempty"`
	EnrolledDescription        string `json:"enrolledDescription,omitempty"`
	InvalidTicketTitle         string `json:"invalidTicketTitle,omitempty"`
	InvalidTicketDescription   string `json:"invalidTicketDescription,omitempty"`
	ExpiredTicketTitle         string `json:"expiredTicketTitle,omitempty"`
	ExpiredTicketDescription   string `json:"expiredTicketDescription,omitempty"`
	AlreadyUsedTitle           string `json:"alreadyUsedTitle,omitempty"`
	AlreadyUsedDescription     string `json:"alreadyUsedDescription,omitempty"`
	AlreadyEnrolledDescription string `json:"alreadyEnrolledDescription,omitempty"`
	GenericError               string `json:"genericError,omitempty"`
}

type ScreenMfaLoginOptions struct {
	PageTitle                          string `json:"pageTitle,omitempty"`
	BackText                           string `json:"backText,omitempty"`
	Title                              string `json:"title,omitempty"`
	AuthenticatorNamesSMS              string `json:"authenticatorNamesSMS,omitempty"`
	AuthenticatorNamesPhone            string `json:"authenticatorNamesPhone,omitempty"`
	AuthenticatorNamesVoice            string `json:"authenticatorNamesVoice,omitempty"`
	AuthenticatorNamesPushNotification string `json:"authenticatorNamesPushNotification,omitempty"`
	AuthenticatorNamesEmail            string `json:"authenticatorNamesEmail,omitempty"`
	AuthenticatorNamesRecoveryCode     string `json:"authenticatorNamesRecoveryCode,omitempty"`
	AuthenticatorNamesDUO              string `json:"authenticatorNamesDUO,omitempty"`
	AuthenticatorNamesWebauthnRoaming  string `json:"authenticatorNamesWebauthnRoaming,omitempty"`
	//Documented but not supported
	//AuthenticatorNamesWebauthnPlatform string `json:"authenticatorNamesWebauthnPlatform,omitempty"`
}

type ScreenMfaBeginEnrollOptions struct {
	PageTitle                          string `json:"pageTitle,omitempty"`
	BackText                           string `json:"backText,omitempty"`
	Title                              string `json:"title,omitempty"`
	Description                        string `json:"description,omitempty"`
	LogoAltText                        string `json:"logoAltText,omitempty"`
	AuthenticatorNamesSms              string `json:"authenticatorNamesSMS,omitempty"`
	AuthenticatorNamesVoice            string `json:"authenticatorNamesVoice,omitempty"`
	AuthenticatorNamesPhone            string `json:"authenticatorNamesPhone,omitempty"`
	AuthenticatorNamesPushNotification string `json:"authenticatorNamesPushNotification,omitempty"`
	AuthenticatorNamesOtp              string `json:"authenticatorNamesOTP,omitempty"`
	AuthenticatorNamesEmail            string `json:"authenticatorNamesEmail,omitempty"`
	AuthenticatorNamesRecoveryCode     string `json:"authenticatorNamesRecoveryCode,omitempty"`
	AuthenticatorNamesDUO              string `json:"authenticatorNamesDUO,omitempty"`
	AuthenticatorNamesWebauthnRoaming  string `json:"authenticatorNamesWebauthnRoaming,omitempty"`
	//Documented but not supported
	//AuthenticatorNamesWebauthnPlatform string `json:"authenticatorNamesWebauthnPlatform,omitempty"`
}

type PromptMfa struct {
	MfaEnrollResult       *ScreenMfaEnrollResult       `json:"mfa-enroll-result,omitempty"`
	MfaLoginOptions       *ScreenMfaLoginOptions       `json:"mfa-login-options,omitempty"`
	MfaBeginEnrollOptions *ScreenMfaBeginEnrollOptions `json:"mfa-begin-enroll-options,omitempty"`
}

// ReadPromptMfa retrieves mfa custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfa(language string, opts ...RequestOption) (p *PromptMfa, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfa replaces mfa custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfa(language string, p *PromptMfa, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa", "custom-text", language), p, opts...)
}

// PromptMfaEmail stores mfa-email custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa-email

type ScreenMfaEmailChallenge struct {
	PageTitle                           string `json:"pageTitle,omitempty"`
	BackText                            string `json:"backText,omitempty"`
	ButtonText                          string `json:"buttonText,omitempty"`
	Description                         string `json:"description,omitempty"`
	PickAuthenticatorText               string `json:"pickAuthenticatorText,omitempty"`
	Placeholder                         string `json:"placeholder,omitempty"`
	RememberMeText                      string `json:"rememberMeText,omitempty"`
	ResendActionText                    string `json:"resendActionText,omitempty"`
	ResendText                          string `json:"resendText,omitempty"`
	Title                               string `json:"title,omitempty"`
	InvalidOtpCodeFormat                string `json:"invalid-otp-code-format,omitempty"`
	InvalidCode                         string `json:"invalid-code,omitempty"`
	InvalidExpiredCode                  string `json:"invalid-expired-code,omitempty"`
	AuthenticatorError                  string `json:"authenticator-error,omitempty"`
	NoTransactionInProgress             string `json:"no-transaction-in-progress,omitempty"`
	TooManyEmail                        string `json:"too-many-email,omitempty"`
	TransactionNotFound                 string `json:"transaction-not-found,omitempty"`
	MfaEmailChallengeAuthenticatorError string `json:"mfa-email-challenge-authenticator-error,omitempty"`
}

type ScreenMfaEmailList struct {
	PageTitle string `json:"pageTitle,omitempty"`
	BackText  string `json:"backText,omitempty"`
	Title     string `json:"title,omitempty"`
}

type PromptMfaEmail struct {
	MfaEmailChallenge *ScreenMfaEmailChallenge `json:"mfa-email-challenge,omitempty"`
	MfaEmailList      *ScreenMfaEmailList      `json:"mfa-email-list,omitempty"`
}

// ReadPromptMfaEmail retrieves mfa-email custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfaEmail(language string, opts ...RequestOption) (p *PromptMfaEmail, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa-email", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfaEmail replaces mfa-email custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfaEmail(language string, p *PromptMfaEmail, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa-email", "custom-text", language), p, opts...)
}

// PromptMfaOtp stores mfa-otp custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa-otp

type ScreenMfaOtpEnrollmentQr struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	ButtonText            string `json:"buttonText,omitempty"`
	CodeEnrollmentText    string `json:"codeEnrollmentText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	//Documented but not supported
	//SeparatorText         string `json:"separatorText,omitempty"`
	InvalidOtpCodeFormat string `json:"invalid-otp-code-format,omitempty"`
	InvalidCode          string `json:"invalid-code,omitempty"`
	InvalidExpiredCode   string `json:"invalid-expired-code,omitempty"`
	TooManyFailures      string `json:"too-many-failures,omitempty"`
	TransactionNotFound  string `json:"transaction-not-found,omitempty"`
	UserAlreadyEnrolled  string `json:"user-already-enrolled,omitempty"`
}

type ScreenMfaOtpEnrollmentCode struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	BackText              string `json:"backText,omitempty"`
	ButtonText            string `json:"buttonText,omitempty"`
	AltText               string `json:"altText,omitempty"`
	CopyCodeButtonText    string `json:"copyCodeButtonText,omitempty"`
	Description           string `json:"description,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	Title                 string `json:"title,omitempty"`
	TooManyFailures       string `json:"too-many-failures,omitempty"`
	TransactionNotFound   string `json:"transaction-not-found,omitempty"`
}

type ScreenMfaOtpEnrollmentChallenge struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	ButtonText            string `json:"buttonText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	RememberMeText        string `json:"rememberMeText,omitempty"`
	AuthenticatorError    string `json:"authenticator-error,omitempty"`
	TooManyFailures       string `json:"too-many-failures,omitempty"`
	TransactionNotFound   string `json:"transaction-not-found,omitempty"`
}

type PromptMfaOtp struct {
	MfaOtpEnrollmentQr        *ScreenMfaOtpEnrollmentQr        `json:"mfa-otp-enrollment-qr,omitempty"`
	MfaOtpEnrollmentCode      *ScreenMfaOtpEnrollmentCode      `json:"mfa-otp-enrollment-code,omitempty"`
	MfaOtpEnrollmentChallenge *ScreenMfaOtpEnrollmentChallenge `json:"mfa-otp-challenge,omitempty"`
}

// ReadPromptMfaOtp retrieves mfa-otp custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfaOtp(language string, opts ...RequestOption) (p *PromptMfaOtp, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa-otp", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfaOtp replaces mfa-otp custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfaOtp(language string, p *PromptMfaOtp, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa-otp", "custom-text", language), p, opts...)
}

// PromptMfaPhone stores mfa-phone custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa-phone

type ScreenMfaPhoneChallenge struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	ContinueButtonText    string `json:"continueButtonText,omitempty"`
	SmsButtonText         string `json:"smsButtonText,omitempty"`
	VoiceButtonText       string `json:"voiceButtonText,omitempty"`
	ChooseMessageTypeText string `json:"chooseMessageTypeText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	SendSmsFailed         string `json:"send-sms-failed,omitempty"`
	SendVoiceFailed       string `json:"send-voice-failed,omitempty"`
	InvalidPhoneFormat    string `json:"invalid-phone-format,omitempty"`
	InvalidPhone          string `json:"invalid-phone,omitempty"`
	TooManySms            string `json:"too-many-sms,omitempty"`
	TooManyVoice          string `json:"too-many-voice,omitempty"`
	TransactionNotFound   string `json:"transaction-not-found,omitempty"`
	NoPhone               string `json:"no-phone,omitempty"`
}

type ScreenMfaPhoneEnrollment struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	ContinueButtonText    string `json:"continueButtonText,omitempty"`
	SmsButtonText         string `json:"smsButtonText,omitempty"`
	VoiceButtonText       string `json:"voiceButtonText,omitempty"`
	ChooseMessageTypeText string `json:"chooseMessageTypeText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	SendSmsFailed         string `json:"send-sms-failed,omitempty"`
	SendVoiceFailed       string `json:"send-voice-failed,omitempty"`
	InvalidPhoneFormat    string `json:"invalid-phone-format,omitempty"`
	InvalidPhone          string `json:"invalid-phone,omitempty"`
	TooManySms            string `json:"too-many-sms,omitempty"`
	TooManyVoice          string `json:"too-many-voice,omitempty"`
	TransactionNotFound   string `json:"transaction-not-found,omitempty"`
	NoPhone               string `json:"no-phone,omitempty"`
}

type PromptMfaPhone struct {
	MfaPhoneChallenge  *ScreenMfaPhoneChallenge  `json:"mfa-phone-challenge,omitempty"`
	MfaPhoneEnrollment *ScreenMfaPhoneEnrollment `json:"mfa-phone-enrollment,omitempty"`
}

// ReadPromptMfaPhone retrieves mfa-phone custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfaPhone(language string, opts ...RequestOption) (p *PromptMfaPhone, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa-phone", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfaPhone replaces mfa-phone custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfaPhone(language string, p *PromptMfaPhone, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa-phone", "custom-text", language), p, opts...)
}

// PromptMfaPush stores mfa-push custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa-push

type ScreenMfaPushWelcome struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	AndroidButtonText     string `json:"androidButtonText,omitempty"`
	ButtonText            string `json:"buttonText,omitempty"`
	IosButtonText         string `json:"iosButtonText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
}

type ScreenMfaPushEnrollmentQr struct {
	PageTitle                    string `json:"pageTitle,omitempty"`
	Title                        string `json:"title,omitempty"`
	Description                  string `json:"description,omitempty"`
	PickAuthenticatorText        string `json:"pickAuthenticatorText,omitempty"`
	ButtonText                   string `json:"buttonText,omitempty"`
	EnrollmentTransactionPending string `json:"enrollment-transaction-pending,omitempty"`
}

type ScreenMfaPushChallengePush struct {
	PageTitle   string `json:"pageTitle,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	//Documented but not supported
	//AwaitingConfirmation               string `json:"awaitingConfirmation,omitempty"`
	ButtonText                         string `json:"buttonText,omitempty"`
	PickAuthenticatorText              string `json:"pickAuthenticatorText,omitempty"`
	RememberMeText                     string `json:"rememberMeText,omitempty"`
	ResendActionText                   string `json:"resendActionText,omitempty"`
	ResendText                         string `json:"resendText,omitempty"`
	EnterOtpCode                       string `json:"enterOtpCode,omitempty"`
	SeparatorText                      string `json:"separatorText,omitempty"`
	ChallengeTransactionPending        string `json:"challenge-transaction-pending,omitempty"`
	PollingIntervalExceeded            string `json:"polling-interval-exceeded,omitempty"`
	TooManyPush                        string `json:"too-many-push,omitempty"`
	TransactionNotFound                string `json:"transaction-not-found,omitempty"`
	MfaPushVerifyTransactionPending    string `json:"mfa-push-verify-transaction-pending,omitempty"`
	MfaPushVerifyAuthenticatorError    string `json:"mfa-push-verify-authenticator-error,omitempty"`
	MfaPushChallengeAuthenticatorError string `json:"mfa-push-challenge-authenticator-error,omitempty"`
	TransactionRejected                string `json:"transaction-rejected,omitempty"`
}

type ScreenMfaPushList struct {
	PageTitle string `json:"pageTitle,omitempty"`
	BackText  string `json:"backText,omitempty"`
	Title     string `json:"title,omitempty"`
}

type PromptMfaPush struct {
	MfaPushWelcome       *ScreenMfaPushWelcome       `json:"mfa-push-welcome,omitempty"`
	MfaPushEnrollmentQr  *ScreenMfaPushEnrollmentQr  `json:"mfa-push-enrollment-qr,omitempty"`
	MfaPushChallengePush *ScreenMfaPushChallengePush `json:"mfa-push-challenge-push,omitempty"`
	MfaPushList          *ScreenMfaPushList          `json:"mfa-push-list,omitempty"`
}

// ReadPromptMfaPush retrieves mfa-push custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfaPush(language string, opts ...RequestOption) (p *PromptMfaPush, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa-push", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfaPush replaces mfa-push custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfaPush(language string, p *PromptMfaPush, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa-push", "custom-text", language), p, opts...)
}

// PromptMfaRecoveryCode stores mfa-recovery-code custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa-recovery-code

type ScreenMfaRecoveryCodeEnrollment struct {
	PageTitle          string `json:"pageTitle,omitempty"`
	Title              string `json:"title,omitempty"`
	Description        string `json:"description,omitempty"`
	AltText            string `json:"altText,omitempty"`
	ButtonText         string `json:"buttonText,omitempty"`
	CheckboxText       string `json:"checkboxText,omitempty"`
	CopyCodeButtonText string `json:"copyCodeButtonText,omitempty"`
	NoConfirmation     string `json:"no-confirmation,omitempty"`
}

type ScreenMfaRecoveryCodeChallenge struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	ButtonText            string `json:"buttonText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	InvalidCode           string `json:"invalid-code,omitempty"`
	InvalidCodeFormat     string `json:"invalid-code-format,omitempty"`
	InvalidExpiredCode    string `json:"invalid-expired-code,omitempty"`
	AuthenticatorError    string `json:"authenticator-error,omitempty"`
	NoConfirmation        string `json:"no-confirmation,omitempty"`
	TooManyFailures       string `json:"too-many-failures,omitempty"`
	TransactionNotFound   string `json:"transaction-not-found,omitempty"`
}

type PromptMfaRecoveryCode struct {
	MfaRecoveryCodeEnrollment *ScreenMfaRecoveryCodeEnrollment `json:"mfa-recovery-code-enrollment,omitempty"`
	MfaRecoveryCodeChallenge  *ScreenMfaRecoveryCodeChallenge  `json:"mfa-recovery-code-challenge,omitempty"`
}

// ReadPromptMfaRecoveryCode retrieves mfa-recovery-code custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfaRecoveryCode(language string, opts ...RequestOption) (p *PromptMfaRecoveryCode, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa-recovery-code", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfaRecoveryCode replaces mfa-recovery-code custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfaRecoveryCode(language string, p *PromptMfaRecoveryCode, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa-recovery-code", "custom-text", language), p, opts...)
}

// PromptMfaSms stores mfa-sms custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa-sms

type ScreenMfaCountryCodes struct {
	PageTitle string `json:"pageTitle,omitempty"`
	BackText  string `json:"backText,omitempty"`
	Title     string `json:"title,omitempty"`
}

type ScreenMfaSmsEnrollment struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	ButtonText            string `json:"buttonText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	SendSmsFailed         string `json:"send-sms-failed,omitempty"`
	InvalidPhoneFormat    string `json:"invalid-phone-format,omitempty"`
	InvalidPhone          string `json:"invalid-phone,omitempty"`
	TooManySms            string `json:"too-many-sms,omitempty"`
	TransactionNotFound   string `json:"transaction-not-found,omitempty"`
	NoPhone               string `json:"no-phone,omitempty"`
}

type ScreenMfaSmsChallenge struct {
	PageTitle                            string `json:"pageTitle,omitempty"`
	Title                                string `json:"title,omitempty"`
	Description                          string `json:"description,omitempty"`
	ButtonText                           string `json:"buttonText,omitempty"`
	EditText                             string `json:"editText,omitempty"`
	PickAuthenticatorText                string `json:"pickAuthenticatorText,omitempty"`
	Placeholder                          string `json:"placeholder,omitempty"`
	RememberMeText                       string `json:"rememberMeText,omitempty"`
	ResendActionText                     string `json:"resendActionText,omitempty"`
	ResendText                           string `json:"resendText,omitempty"`
	ResendVoiceActionSeparatorTextBefore string `json:"resendVoiceActionSeparatorTextBefore,omitempty"`
	ResendVoiceActionText                string `json:"resendVoiceActionText,omitempty"`
	ResendVoiceActionSeparatorTextAfter  string `json:"resendVoiceActionSeparatorTextAfter,omitempty"`
	InvalidOtpCodeFormat                 string `json:"invalid-otp-code-format,omitempty"`
	InvalidCode                          string `json:"invalid-code,omitempty"`
	InvalidExpiredCode                   string `json:"invalid-expired-code,omitempty"`
	SendSmsFailed                        string `json:"send-sms-failed,omitempty"`
	AuthenticatorError                   string `json:"authenticator-error,omitempty"`
	SmsAuthenticatorError                string `json:"sms-authenticator-error,omitempty"`
	NoTransactionInProgress              string `json:"no-transaction-in-progress,omitempty"`
	TooManyFailures                      string `json:"too-many-failures,omitempty"`
	TooManySms                           string `json:"too-many-sms,omitempty"`
	TransactionNotFound                  string `json:"transaction-not-found,omitempty"`
}

type ScreenMfaSmsList struct {
	PageTitle string `json:"pageTitle,omitempty"`
	BackText  string `json:"backText,omitempty"`
	Title     string `json:"title,omitempty"`
}

type PromptMfaSms struct {
	MfaCountryCodes  *ScreenMfaCountryCodes  `json:"mfa-country-codes,omitempty"`
	MfaSmsEnrollment *ScreenMfaSmsEnrollment `json:"mfa-sms-enrollment,omitempty"`
	MfaSmsChallenge  *ScreenMfaSmsChallenge  `json:"mfa-sms-challenge,omitempty"`
	MfaSmsList       *ScreenMfaSmsList       `json:"mfa-sms-list,omitempty"`
}

// ReadPromptMfaSms retrieves mfa-sms custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfaSms(language string, opts ...RequestOption) (p *PromptMfaSms, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa-sms", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfaSms replaces mfa-sms custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfaSms(language string, p *PromptMfaSms, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa-sms", "custom-text", language), p, opts...)
}

// PromptMfaVoice store mfa-voice custom text
//
// See: https://auth0.com/docs/universal-login/prompt-mfa-voice

type ScreenMfaVoiceEnrollment struct {
	PageTitle             string `json:"pageTitle,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	ButtonText            string `json:"buttonText,omitempty"`
	PickAuthenticatorText string `json:"pickAuthenticatorText,omitempty"`
	Placeholder           string `json:"placeholder,omitempty"`
	SendSmsFailed         string `json:"send-sms-failed,omitempty"`
	InvalidPhoneFormat    string `json:"invalid-phone-format,omitempty"`
	InvalidPhone          string `json:"invalid-phone,omitempty"`
	TooManySms            string `json:"too-many-sms,omitempty"`
	TransactionNotFound   string `json:"transaction-not-found,omitempty"`
	NoPhone               string `json:"no-phone,omitempty"`
}

type ScreenMfaVoiceChallenge struct {
	PageTitle                          string `json:"pageTitle,omitempty"`
	Title                              string `json:"title,omitempty"`
	Description                        string `json:"description,omitempty"`
	ButtonText                         string `json:"buttonText,omitempty"`
	EditText                           string `json:"editText,omitempty"`
	PickAuthenticatorText              string `json:"pickAuthenticatorText,omitempty"`
	Placeholder                        string `json:"placeholder,omitempty"`
	RememberMeText                     string `json:"rememberMeText,omitempty"`
	ResendActionText                   string `json:"resendActionText,omitempty"`
	ResendText                         string `json:"resendText,omitempty"`
	ResendSmsActionSeparatorTextBefore string `json:"resendSmsActionSeparatorTextBefore,omitempty"`
	ResendSmsActionText                string `json:"resendSmsActionText,omitempty"`
	ResendSmsActionSeparatorTextAfter  string `json:"resendSmsActionSeparatorTextAfter,omitempty"`
	InvalidOtpCodeFormat               string `json:"invalid-otp-code-format,omitempty"`
	InvalidCode                        string `json:"invalid-code,omitempty"`
	InvalidExpiredCode                 string `json:"invalid-expired-code,omitempty"`
	SendVoiceFailed                    string `json:"send-voice-failed,omitempty"`
	AuthenticatorError                 string `json:"authenticator-error,omitempty"`
	VoiceAuthenticatorError            string `json:"voice-authenticator-error,omitempty"`
	NoTransactionInProgress            string `json:"no-transaction-in-progress,omitempty"`
	TooManyFailures                    string `json:"too-many-failures,omitempty"`
	TooManyVoice                       string `json:"too-many-voice,omitempty"`
	TransactionNotFound                string `json:"transaction-not-found,omitempty"`
	NoPhone                            string `json:"no-phone,omitempty"`
}

type PromptMfaVoice struct {
	MfaVoiceEnrollment *ScreenMfaVoiceEnrollment `json:"mfa-voice-enrollment,omitempty"`
	MfaVoiceChallenge  *ScreenMfaVoiceChallenge  `json:"mfa-voice-challenge,omitempty"`
}

// ReadPromptMfaVoice retrieves mfa-voice custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptMfaVoice(language string, opts ...RequestOption) (p *PromptMfaVoice, err error) {
	err = m.Request("GET", m.URI("prompts", "mfa-voice", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptMfaVoice replaces mfa-voice custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptMfaVoice(language string, p *PromptMfaVoice, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "mfa-voice", "custom-text", language), p, opts...)
}

// PromptOrganizationSelection store organization-selection custom text
//
// See: https://auth0.com/docs/universal-login/prompt-organization-selection

type ScreenOrganizationSelection struct {
	PageTitle   string `json:"pageTitle,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ButtonText  string `json:"buttonText,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	//Documented but not supported
	//ErrorInvalidOrganization string `json:"error_invalid-organization,omitempty"`
	LogoAltText string `json:"logoAltText,omitempty"`
}

type PromptOrganizationSelection struct {
	OrganizationSelection *ScreenOrganizationSelection `json:"organization-selection,omitempty"`
}

// ReadPromptOrganizations retrieves organizations custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptOrganizations(language string, opts ...RequestOption) (p *PromptOrganizationSelection, err error) {
	err = m.Request("GET", m.URI("prompts", "organizations", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptOrganizations replaces organizations custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptOrganizations(language string, p *PromptOrganizationSelection, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "organizations", "custom-text", language), p, opts...)
}

// PromptResetPassword stores reset-password custom text
//
// See: https://auth0.com/docs/universal-login/prompt-reset-password

type ScreenResetPasswordRequest struct {
	PageTitle               string `json:"pageTitle,omitempty"`
	Title                   string `json:"title,omitempty"`
	BackToLoginLinkText     string `json:"backToLoginLinkText,omitempty"`
	ButtonText              string `json:"buttonText,omitempty"`
	DescriptionEmail        string `json:"descriptionEmail,omitempty"`
	DescriptionUsername     string `json:"descriptionUsername,omitempty"`
	PlaceholderEmail        string `json:"placeholderEmail,omitempty"`
	PlaceholderUsername     string `json:"placeholderUsername,omitempty"`
	InvalidEmailFormat      string `json:"invalid-email-format,omitempty"`
	Auth0UsersExpiredTicket string `json:"auth0-users-expired-ticket,omitempty"`
	CustomScriptErrorCode   string `json:"custom-script-error-code,omitempty"`
	Auth0UsersUsedTicket    string `json:"auth0-users-used-ticket,omitempty"`
	Auth0UsersValidation    string `json:"auth0-users-validation,omitempty"`
	ResetPasswordError      string `json:"reset-password-error,omitempty"`
	TooManyEmail            string `json:"too-many-email,omitempty"`
	TooManyRequests         string `json:"too-many-requests,omitempty"`
	NoEmail                 string `json:"no-email,omitempty"`
	NoUsername              string `json:"no-username,omitempty"`
}

type ScreenResetPasswordEmail struct {
	PageTitle        string `json:"pageTitle,omitempty"`
	Title            string `json:"title,omitempty"`
	EmailDescription string `json:"emailDescription,omitempty"`
	ResendLinkText   string `json:"resendLinkText,omitempty"`
	//Documented but not supported
	//ResendText          string `json:"resendText,omitempty"`
	UsernameDescription string `json:"usernameDescription,omitempty"`
}

type ScreenResetPassword struct {
	PageTitle                  string `json:"pageTitle,omitempty"`
	Title                      string `json:"title,omitempty"`
	Description                string `json:"description,omitempty"`
	ButtonText                 string `json:"buttonText,omitempty"`
	PasswordPlaceholder        string `json:"passwordPlaceholder,omitempty"`
	ReEnterPasswordPlaceholder string `json:"reEnterpasswordPlaceholder,omitempty"`
	PasswordSecurityText       string `json:"passwordSecurityText,omitempty"`
	Auth0UsersExpiredTicket    string `json:"auth0-users-expired-ticket,omitempty"`
	CustomScriptErrorCode      string `json:"custom-script-error-code,omitempty"`
	Auth0UsersUsedTicket       string `json:"auth0-users-used-ticket,omitempty"`
	Auth0UsersValidation       string `json:"auth0-users-validation,omitempty"`
	NoReEnterPassword          string `json:"no-re-enter-password,omitempty"`
}

type ScreenResetPasswordSuccess struct {
	PageTitle   string `json:"pageTitle,omitempty"`
	EventTitle  string `json:"eventTitle,omitempty"`
	Description string `json:"description,omitempty"`
	ButtonText  string `json:"buttonText,omitempty"`
}

type ScreenResetPasswordError struct {
	PageTitle               string `json:"pageTitle,omitempty"`
	BackToLoginLinkText     string `json:"backToLoginLinkText,omitempty"`
	DescriptionExpired      string `json:"descriptionExpired,omitempty"`
	DescriptionGeneric      string `json:"descriptionGeneric,omitempty"`
	DescriptionUsed         string `json:"descriptionUsed,omitempty"`
	EventTitleExpired       string `json:"eventTitleExpired,omitempty"`
	EventTitleUsed          string `json:"eventTitleUsed,omitempty"`
	Auth0UsersExpiredTicket string `json:"auth0-users-expired-ticket,omitempty"`
	CustomScriptErrorCode   string `json:"custom-script-error-code,omitempty"`
	Auth0UsersUsedTicket    string `json:"auth0-users-used-ticket,omitempty"`
	Auth0UsersValidation    string `json:"auth0-users-validation,omitempty"`
	ResetPasswordError      string `json:"reset-password-error,omitempty"`
}

type PromptResetPassword struct {
	ResetPasswordRequest *ScreenResetPasswordRequest `json:"reset-password-request,omitempty"`
	ResetPasswordEmail   *ScreenResetPasswordEmail   `json:"reset-password-email,omitempty"`
	ResetPassword        *ScreenResetPassword        `json:"reset-password,omitempty"`
	ResetPasswordSuccess *ScreenResetPasswordSuccess `json:"reset-password-success,omitempty"`
	ResetPasswordError   *ScreenResetPasswordError   `json:"reset-password-error,omitempty"`
}

// ReadPromptResetPassword retrieves reset-password custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptResetPassword(language string, opts ...RequestOption) (p *PromptResetPassword, err error) {
	err = m.Request("GET", m.URI("prompts", "reset-password", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptResetPassword replaces reset-password custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptResetPassword(language string, p *PromptResetPassword, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "reset-password", "custom-text", language), p, opts...)
}

// PromptSignup stores signup custom text
//
// See: https://auth0.com/docs/universal-login/prompt-signup

type ScreenSignup struct {
	PageTitle                        string `json:"pageTitle,omitempty"`
	Title                            string `json:"title,omitempty"`
	Description                      string `json:"description,omitempty"`
	SeparatorText                    string `json:"separatorText,omitempty"`
	ButtonText                       string `json:"buttonText,omitempty"`
	EmailPlaceholder                 string `json:"emailPlaceholder,omitempty"`
	FederatedConnectionButtonText    string `json:"federatedConnectionButtonText,omitempty"`
	LoginActionLinkText              string `json:"loginActionLinkText,omitempty"`
	LoginActionText                  string `json:"loginActionText,omitempty"`
	PasswordPlaceholder              string `json:"passwordPlaceholder,omitempty"`
	PasswordSecurityText             string `json:"passwordSecurityText,omitempty"`
	UsernamePlaceholder              string `json:"usernamePlaceholder,omitempty"`
	EmailInUse                       string `json:"email-in-use,omitempty"`
	InvalidEmailFormat               string `json:"invalid-email-format,omitempty"`
	PasswordTooWeak                  string `json:"password-too-common,omitempty"`
	PasswordTooCommon                string `json:"password-too-weak,omitempty"`
	PasswordPreviouslyUsed           string `json:"password-previously-used,omitempty"`
	PasswordMismatch                 string `json:"password-mismatch,omitempty"`
	InvalidUsername                  string `json:"invalid-username,omitempty"`
	InvalidUsernameMaxLength         string `json:"invalid-username-max-length,omitempty"`
	InvalidUsernameMinLength         string `json:"invalid-username-min-length,omitempty"`
	InvalidUsernameInvalidCharacters string `json:"invalid-username-invalid-characters,omitempty"`
	InvalidUsernameEmailNotAllowed   string `json:"invalid-username-email-not-allowed,omitempty"`
	UsernameTaken                    string `json:"username-taken,omitempty"`
	CustomScriptErrorCode            string `json:"custom-script-error-code,omitempty"`
	Auth0UsersValidation             string `json:"auth0-users-validation,omitempty"`
	InvalidConnection                string `json:"invalid-connection,omitempty"`
	IpBlocked                        string `json:"ip-blocked,omitempty"`
	IpSignupBlocked                  string `json:"ip-signup-blocked,omitempty"`
	NoDbConnection                   string `json:"no-db-connection,omitempty"`
	NoEmail                          string `json:"no-email,omitempty"`
	NoPassword                       string `json:"no-password,omitempty"`
	NoReEnterPassword                string `json:"no-re-enter-password,omitempty"`
	NoUsername                       string `json:"no-username,omitempty"`
}

type PromptSignup struct {
	Signup *ScreenSignup `json:"signup,omitempty"`
}

// ReadPromptSignup retrieves signup custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptSignup(language string, opts ...RequestOption) (p *PromptSignup, err error) {
	err = m.Request("GET", m.URI("prompts", "signup", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptSignup replaces signup custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptSignup(language string, p *PromptSignup, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "signup", "custom-text", language), p, opts...)
}

// PromptSignupId store signup-id custom text
//
// See: https://auth0.com/docs/universal-login/prompt-signup-id

type ScreenSignupId struct {
	PageTitle                        string `json:"pageTitle,omitempty"`
	Title                            string `json:"title,omitempty"`
	Description                      string `json:"description,omitempty"`
	SeparatorText                    string `json:"separatorText,omitempty"`
	ButtonText                       string `json:"buttonText,omitempty"`
	EmailPlaceholder                 string `json:"emailPlaceholder,omitempty"`
	FederatedConnectionButtonText    string `json:"federatedConnectionButtonText,omitempty"`
	LoginActionLinkText              string `json:"loginActionLinkText,omitempty"`
	LoginActionText                  string `json:"loginActionText,omitempty"`
	PasswordPlaceholder              string `json:"passwordPlaceholder,omitempty"`
	PasswordSecurityText             string `json:"passwordSecurityText,omitempty"`
	UsernamePlaceholder              string `json:"usernamePlaceholder,omitempty"`
	LogoAltText                      string `json:"logoAltText,omitempty"`
	EmailInUse                       string `json:"email-in-use,omitempty"`
	InvalidEmailFormat               string `json:"invalid-email-format,omitempty"`
	PasswordTooWeak                  string `json:"password-too-common,omitempty"`
	PasswordTooCommon                string `json:"password-too-weak,omitempty"`
	PasswordPreviouslyUsed           string `json:"password-previously-used,omitempty"`
	PasswordMismatch                 string `json:"password-mismatch,omitempty"`
	InvalidUsername                  string `json:"invalid-username,omitempty"`
	InvalidUsernameMaxLength         string `json:"invalid-username-max-length,omitempty"`
	InvalidUsernameMinLength         string `json:"invalid-username-min-length,omitempty"`
	InvalidUsernameInvalidCharacters string `json:"invalid-username-invalid-characters,omitempty"`
	InvalidUsernameEmailNotAllowed   string `json:"invalid-username-email-not-allowed,omitempty"`
	UsernameTaken                    string `json:"username-taken,omitempty"`
	CustomScriptErrorCode            string `json:"custom-script-error-code,omitempty"`
	Auth0UsersValidation             string `json:"auth0-users-validation,omitempty"`
	InvalidConnection                string `json:"invalid-connection,omitempty"`
	IpBlocked                        string `json:"ip-blocked,omitempty"`
	IpSignupBlocked                  string `json:"ip-signup-blocked,omitempty"`
	NoDbConnection                   string `json:"no-db-connection,omitempty"`
	NoEmail                          string `json:"no-email,omitempty"`
	NoPassword                       string `json:"no-password,omitempty"`
	NoReEnterPassword                string `json:"no-re-enter-password,omitempty"`
	NoUsername                       string `json:"no-username,omitempty"`
}

type PromptSignupId struct {
	SignupId *ScreenSignupId `json:"signup-id,omitempty"`
}

// ReadPromptSignupId retrieves signup-id custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptSignupId(language string, opts ...RequestOption) (p *PromptSignupId, err error) {
	err = m.Request("GET", m.URI("prompts", "signup-id", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptSignupId replaces signup-id custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptSignupId(language string, p *PromptSignupId, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "signup-id", "custom-text", language), p, opts...)
}

// PromptSignupPassword stores signup-password custom text
//
// See: https://auth0.com/docs/universal-login/prompt-signup-password

type ScreenSignupPassword struct {
	PageTitle                     string `json:"pageTitle,omitempty"`
	Title                         string `json:"title,omitempty"`
	Description                   string `json:"description,omitempty"`
	SeparatorText                 string `json:"separatorText,omitempty"`
	ButtonText                    string `json:"buttonText,omitempty"`
	EmailPlaceholder              string `json:"emailPlaceholder,omitempty"`
	FederatedConnectionButtonText string `json:"federatedConnectionButtonText,omitempty"`
	LoginActionLinkText           string `json:"loginActionLinkText,omitempty"`
	LoginActionText               string `json:"loginActionText,omitempty"`
	//Documented but not supported
	//LoginLinkText                    string `json:"loginLinkText,omitempty"`
	InvitationTitle                  string `json:"invitationTitle,omitempty"`
	InvitationDescription            string `json:"invitationDescription,omitempty"`
	PasswordPlaceholder              string `json:"passwordPlaceholder,omitempty"`
	PasswordSecurityText             string `json:"passwordSecurityText,omitempty"`
	UsernamePlaceholder              string `json:"usernamePlaceholder,omitempty"`
	EmailInUse                       string `json:"email-in-use,omitempty"`
	InvalidEmailFormat               string `json:"invalid-email-format,omitempty"`
	PasswordTooWeak                  string `json:"password-too-common,omitempty"`
	PasswordTooCommon                string `json:"password-too-weak,omitempty"`
	PasswordPreviouslyUsed           string `json:"password-previously-used,omitempty"`
	PasswordMismatch                 string `json:"password-mismatch,omitempty"`
	InvalidUsername                  string `json:"invalid-username,omitempty"`
	InvalidUsernameMaxLength         string `json:"invalid-username-max-length,omitempty"`
	InvalidUsernameMinLength         string `json:"invalid-username-min-length,omitempty"`
	InvalidUsernameInvalidCharacters string `json:"invalid-username-invalid-characters,omitempty"`
	InvalidUsernameEmailNotAllowed   string `json:"invalid-username-email-not-allowed,omitempty"`
	UsernameTaken                    string `json:"username-taken,omitempty"`
	CustomScriptErrorCode            string `json:"custom-script-error-code,omitempty"`
	Auth0UsersValidation             string `json:"auth0-users-validation,omitempty"`
	InvalidConnection                string `json:"invalid-connection,omitempty"`
	IpBlocked                        string `json:"ip-blocked,omitempty"`
	IpSignupBlocked                  string `json:"ip-signup-blocked,omitempty"`
	NoDbConnection                   string `json:"no-db-connection,omitempty"`
	NoEmail                          string `json:"no-email,omitempty"`
	NoPassword                       string `json:"no-password,omitempty"`
	NoReEnterPassword                string `json:"no-re-enter-password,omitempty"`
	NoUsername                       string `json:"no-username,omitempty"`
	LogoAltText                      string `json:"logoAltText,omitempty"`
}

type PromptSignupPassword struct {
	SignupPassword *ScreenSignupPassword `json:"signup-password,omitempty"`
}

// ReadPromptSignupPassword retrieves signup-password custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) ReadPromptSignupPassword(language string, opts ...RequestOption) (p *PromptSignupPassword, err error) {
	err = m.Request("GET", m.URI("prompts", "signup-password", "custom-text", language), &p, opts...)
	return
}

// ReplacePromptSignupPassword replaces signup-password custom text
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) ReplacePromptSignupPassword(language string, p *PromptSignupPassword, opts ...RequestOption) error {
	return m.Request("PUT", m.URI("prompts", "signup-password", "custom-text", language), p, opts...)
}
