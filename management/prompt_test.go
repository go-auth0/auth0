package management

import (
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/internal/testing/expect"
	"testing"
)

func TestPrompt(t *testing.T) {
	t.Cleanup(func() {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "classic",
			IdentifierFirst:          auth0.Bool(false),
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	// update to the new identifier first experience
	err := m.Prompt.Update(&Prompt{
		UniversalLoginExperience: "new",
		IdentifierFirst:          auth0.Bool(true),
	})
	if err != nil {
		t.Error(err)
	}

	ps, err := m.Prompt.Read()
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, ps.UniversalLoginExperience, "new")
	expect.Expect(t, ps.IdentifierFirst, auth0.Bool(true))

	// update to the classic non identifier first experience
	err = m.Prompt.Update(&Prompt{
		UniversalLoginExperience: "classic",
		IdentifierFirst:          auth0.Bool(false),
	})
	if err != nil {
		t.Error(err)
	}

	ps, err = m.Prompt.Read()
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, ps.UniversalLoginExperience, "classic")
	expect.Expect(t, ps.IdentifierFirst, auth0.Bool(false))
}

// Language to use with tests
const Language = "en"

func TestPromptConsent(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptConsent(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptConsent(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptConsent(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptConsent{
		Consent: &ScreenConsent{
			PageTitle:              "PageTitle",
			Title:                  "Title",
			MessageMultipleTenants: "MessageMultipleTenants",
			AudiencePickerAltText:  "AudiencePickerAltText",
			MessageSingleTenant:    "MessageSingleTenant",
			AcceptButtonText:       "AcceptButtonText",
			DeclineButtonText:      "DeclineButtonText",
			InvalidAction:          "InvalidAction",
			InvalidAudience:        "InvalidAudience",
			InvalidScope:           "InvalidScope",
		},
	}

	err = m.Prompt.ReplacePromptConsent(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptConsent(Language)
	if err != nil {
		t.Fatal(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptDeviceFlow(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptDeviceFlow(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptDeviceFlow(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptDeviceFlow(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptDeviceFlow{
		DeviceCodeActivation: &ScreenDeviceCodeActivation{
			PageTitle:          "PageTitle",
			ButtonText:         "ButtonText",
			Description:        "Description",
			Placeholder:        "Placeholder",
			Title:              "Title",
			InvalidExpiredCode: "InvalidExpiredCode",
			NoCode:             "NoCode",
			InvalidCode:        "InvalidCode",
		},
		DeviceCodeActivationAllowed: &ScreenDeviceCodeActivationAllowed{
			PageTitle:   "PageTitle",
			Description: "Description",
			EventTitle:  "EventTitle",
		},
		DeviceCodeActivationDenied: &ScreenDeviceCodeActivationDenied{
			PageTitle:   "PageTitle",
			Description: "Description",
			EventTitle:  "EventTitle",
		},
		DeviceCodeConfirmation: &ScreenDeviceCodeConfirmation{
			PageTitle:         "PageTitle",
			Description:       "Description",
			InputCodeLabel:    "InputCodeLabel",
			Title:             "Title",
			ConfirmButtonText: "ConfirmButtonText",
			CancelButtonText:  "CancelButtonText",
			ConfirmationText:  "ConfirmationText",
		},
	}

	err = m.Prompt.ReplacePromptDeviceFlow(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptDeviceFlow(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptEmailOtpChallenge(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptEmailOtpChallenge(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptEmailOtpChallenge(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptEmailOtpChallenge(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptEmailOtpChallenge{
		EmailOtpChallenge: &ScreenEmailOtpChallenge{
			PageTitle:            "PageTitle",
			ButtonText:           "ButtonText",
			Description:          "Description",
			Placeholder:          "Placeholder",
			ResendActionText:     "ResendActionText",
			ResendText:           "ResendText",
			Title:                "Title",
			InvalidOtpCodeFormat: "InvalidOtpCodeFormat",
			InvalidCode:          "InvalidCode",
			InvalidExpiredCode:   "InvalidExpiredCode",
			AuthenticatorError:   "AuthenticatorError",
			TooManyEmail:         "TooManyEmail",
		},
	}

	err = m.Prompt.ReplacePromptEmailOtpChallenge(Language, &SetVal)
	if err != nil {
		t.Fatal(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptEmailOtpChallenge(Language)
	if err != nil {
		t.Fatal(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptEmailVerification(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptEmailVerification(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptEmailVerification(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptEmailVerification(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptEmailVerification{
		EmailVerificationResult: &ScreenEmailVerificationResult{
			PageTitle:                       "PageTitle",
			VerifiedTitle:                   "VerifiedTitle",
			ErrorTitle:                      "ErrorTitle",
			VerifiedDescription:             "VerifiedDescription",
			AlreadyVerifiedDescription:      "AlreadyVerifiedDescription",
			InvalidAccountOrCodeDescription: "InvalidAccountOrCodeDescription",
			UnknownErrorDescription:         "UnknownErrorDescription",
			ButtonText:                      "ButtonText",
			Auth0UsersExpiredTicket:         "Auth0UsersExpiredTicket",
			CustomScriptErrorCode:           "CustomScriptErrorCode",
			Auth0UsersUsedTicket:            "Auth0UsersUsedTicket",
			Auth0UsersValidation:            "Auth0UsersValidation",
		},
	}

	err = m.Prompt.ReplacePromptEmailVerification(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptEmailVerification(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptAcceptInvitation(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptAcceptInvitation(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptAcceptInvitation(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptAcceptInvitation(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptAcceptInvitation{
		AcceptInvitation: &ScreenAcceptInvitation{
			PageTitle:   "PageTitle",
			Title:       "Title",
			Description: "Description",
			ButtonText:  "ButtonText",
			LogoAltText: "LogoAltText",
		},
	}

	err = m.Prompt.ReplacePromptAcceptInvitation(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptAcceptInvitation(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptLogin(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptLogin(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptLogin(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptLogin(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptLogin{
		Login: &ScreenLogin{
			PageTitle:                     "PageTitle",
			Title:                         "Title",
			Description:                   "Description",
			SeparatorText:                 "SeparatorText",
			ButtonText:                    "ButtonText",
			FederatedConnectionButtonText: "FederatedConnectionButtonText",
			SignupActionLinkText:          "SignupActionLinkText",
			SignupActionText:              "SignupActionTextc",
			ForgotPasswordText:            "ForgotPasswordText",
			PasswordPlaceholder:           "PasswordPlaceholder",
			UsernamePlaceholder:           "UsernamePlaceholder",
			//In online documentation but not supported
			//CaptchaCodePlaceholder: "CaptchaCodePlaceholder",
			//CaptchaMatchExprPlaceholder: "CaptchaMatchExprPlaceholder",
			EmailPlaceholder:      "EmailPlaceholder",
			EditEmailText:         "EditEmailText",
			AlertListTitle:        "AlertListTitle",
			InvitationTitle:       "InvitationTitle",
			InvitationDescription: "InvitationDescription",
			WrongCredentials:      "WrongCredentials",
			//In online documentation but not supported
			//WrongCaptcha: "WrongCaptcha",
			InvalidCode:           "InvalidCode",
			InvalidExpiredCode:    "InvalidExpiredCode",
			InvalidEmailFormat:    "InvalidEmailFormat",
			WrongEmailCredentials: "WrongEmailCredentials",
			CustomScriptErrorCode: "CustomScriptErrorCode",
			Auth0UsersValidation:  "Auth0UsersValidation",
			AuthenticationFailure: "AuthenticationFailure",
			InvalidConnection:     "InvalidConnection",
			IpBlocked:             "IpBlocked",
			NoDbConnection:        "NoDbConnection",
			PasswordBreached:      "PasswordBreached",
			UserBlocked:           "UserBlocked",
			SameUserLogin:         "SameUserLogin",
			NoEmail:               "NoEmail",
			NoPassword:            "NoPassword",
			NoUsername:            "NoUsername",
		},
	}

	err = m.Prompt.ReplacePromptLogin(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptLogin(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptLoginId(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptLoginId(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptLoginId(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptLoginId(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptLoginId{
		LoginId: &ScreenLoginId{
			PageTitle:                     "PageTitle",
			Title:                         "Title",
			Description:                   "Description",
			SeparatorText:                 "SeparatorText",
			ButtonText:                    "ButtonText",
			FederatedConnectionButtonText: "FederatedConnectionButtonText",
			SignupActionLinkText:          "SignupActionLinkText",
			SignupActionText:              "SignupActionTextc",
			PasswordPlaceholder:           "PasswordPlaceholder",
			UsernamePlaceholder:           "UsernamePlaceholder",
			EmailPlaceholder:              "EmailPlaceholder",
			EditEmailText:                 "EditEmailText",
			AlertListTitle:                "AlertListTitle",
			LogoAltText:                   "LogoAltText",
			WrongCredentials:              "WrongCredentials",
			InvalidCode:                   "InvalidCode",
			InvalidExpiredCode:            "InvalidExpiredCode",
			InvalidEmailFormat:            "InvalidEmailFormat",
			WrongEmailCredentials:         "WrongEmailCredentials",
			CustomScriptErrorCode:         "CustomScriptErrorCode",
			Auth0UsersValidation:          "Auth0UsersValidation",
			AuthenticationFailure:         "AuthenticationFailure",
			InvalidConnection:             "InvalidConnection",
			IpBlocked:                     "IpBlocked",
			NoDbConnection:                "NoDbConnection",
			NoHrdConnection:               "NoHrdConnection",
			PasswordBreached:              "PasswordBreached",
			UserBlocked:                   "UserBlocked",
			SameUserLogin:                 "SameUserLogin",
			NoEmail:                       "NoEmail",
			NoPassword:                    "NoPassword",
			NoUsername:                    "NoUsername",
		},
	}

	err = m.Prompt.ReplacePromptLoginId(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptLoginId(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptLoginPassword(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptLoginPassword(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptLoginPassword(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptLoginPassword(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptLoginPassword{
		LoginPassword: &ScreenLoginPassword{
			PageTitle:                     "PageTitle",
			Title:                         "Title",
			Description:                   "Description",
			SeparatorText:                 "SeparatorText",
			ButtonText:                    "ButtonText",
			FederatedConnectionButtonText: "FederatedConnectionButtonText",
			SignupActionLinkText:          "SignupActionLinkText",
			SignupActionText:              "SignupActionTextc",
			PasswordPlaceholder:           "PasswordPlaceholder",
			UsernamePlaceholder:           "UsernamePlaceholder",
			EmailPlaceholder:              "EmailPlaceholder",
			EditEmailText:                 "EditEmailText",
			AlertListTitle:                "AlertListTitle",
			LogoAltText:                   "LogoAltText",
			WrongCredentials:              "WrongCredentials",
			InvalidCode:                   "InvalidCode",
			InvalidExpiredCode:            "InvalidExpiredCode",
			InvalidEmailFormat:            "InvalidEmailFormat",
			WrongEmailCredentials:         "WrongEmailCredentials",
			CustomScriptErrorCode:         "CustomScriptErrorCode",
			Auth0UsersValidation:          "Auth0UsersValidation",
			AuthenticationFailure:         "AuthenticationFailure",
			InvalidConnection:             "InvalidConnection",
			IpBlocked:                     "IpBlocked",
			NoDbConnection:                "NoDbConnection",
			PasswordBreached:              "PasswordBreached",
			UserBlocked:                   "UserBlocked",
			SameUserLogin:                 "SameUserLogin",
			NoEmail:                       "NoEmail",
			NoPassword:                    "NoPassword",
			NoUsername:                    "NoUsername",
		},
	}

	err = m.Prompt.ReplacePromptLoginPassword(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptLoginPassword(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptLoginEmailVerification(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptLoginEmailVerification(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptLoginEmailVerification(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptLoginEmailVerification(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptLoginEmailVerification{
		LoginEmailVerification: &ScreenLoginEmailVerification{
			PageTitle:            "PageTitle",
			ButtonText:           "ButtonText",
			Description:          "Description",
			Placeholder:          "Placeholder",
			ResendActionText:     "ResendActionText",
			ResendText:           "ResendText",
			Title:                "Title",
			InvalidOtpCodeFormat: "InvalidOtpCodeFormat",
			InvalidCode:          "InvalidCode",
			InvalidExpiredCode:   "InvalidExpiredCode",
			AuthenticatorError:   "AuthenticatorError",
			TooManyEmail:         "TooManyEmail",
		},
	}

	err = m.Prompt.ReplacePromptLoginEmailVerification(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptLoginEmailVerification(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfa(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfa(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfa(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfa(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfa{
		MfaEnrollResult: &ScreenMfaEnrollResult{
			PageTitle:                  "PageTitle",
			EnrolledTitle:              "EnrolledTitle",
			EnrolledDescription:        "EnrolledDescription",
			InvalidTicketTitle:         "InvalidTicketTitle",
			InvalidTicketDescription:   "InvalidTicketDescription",
			ExpiredTicketTitle:         "ExpiredTicketTitle",
			ExpiredTicketDescription:   "ExpiredTicketDescription",
			AlreadyUsedTitle:           "AlreadyUsedTitle",
			AlreadyUsedDescription:     "AlreadyUsedDescription",
			AlreadyEnrolledDescription: "AlreadyEnrolledDescription",
			GenericError:               "GenericError",
		},
		MfaLoginOptions: &ScreenMfaLoginOptions{
			PageTitle:                          "PageTitle",
			BackText:                           "BackText",
			Title:                              "Title",
			AuthenticatorNamesSMS:              "AuthenticatorNamesSMS",
			AuthenticatorNamesPhone:            "AuthenticatorNamesPhone",
			AuthenticatorNamesVoice:            "AuthenticatorNamesVoice",
			AuthenticatorNamesPushNotification: "AuthenticatorNamesPushNotification",
			AuthenticatorNamesEmail:            "AuthenticatorNamesEmail",
			AuthenticatorNamesRecoveryCode:     "AuthenticatorNamesRecoveryCode",
			AuthenticatorNamesDUO:              "AuthenticatorNamesDUO",
			AuthenticatorNamesWebauthnRoaming:  "AuthenticatorNamesWebauthnRoaming",
			//In online documentation but not supported
			// AuthenticatorNamesWebauthnPlatform: "authenticatorNamesWebauthnPlatform",
		},
		MfaBeginEnrollOptions: &ScreenMfaBeginEnrollOptions{
			PageTitle:                          "PageTitle",
			BackText:                           "BackText",
			Title:                              "Title",
			Description:                        "Description",
			LogoAltText:                        "LogoAltText",
			AuthenticatorNamesSms:              "AuthenticatorNamesSms",
			AuthenticatorNamesVoice:            "AuthenticatorNamesVoice",
			AuthenticatorNamesPhone:            "AuthenticatorNamesPhone",
			AuthenticatorNamesPushNotification: "AuthenticatorNamesPushNotification",
			AuthenticatorNamesOtp:              "AuthenticatorNamesOtp",
			AuthenticatorNamesEmail:            "AuthenticatorNamesEmail",
			AuthenticatorNamesRecoveryCode:     "AuthenticatorNamesRecoveryCode",
			AuthenticatorNamesDUO:              "AuthenticatorNamesDUO",
			AuthenticatorNamesWebauthnRoaming:  "AuthenticatorNamesWebauthnRoaming",
			//In online documentation but not supported
			// AuthenticatorNamesWebauthnPlatform: "authenticatorNamesWebauthnPlatform",
		},
	}

	err = m.Prompt.ReplacePromptMfa(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfa(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfaEmail(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfaEmail(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfaEmail(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfaEmail(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfaEmail{
		MfaEmailChallenge: &ScreenMfaEmailChallenge{
			PageTitle:                           "PageTitle",
			BackText:                            "BackText",
			ButtonText:                          "ButtonText",
			Description:                         "Description",
			PickAuthenticatorText:               "PickAuthenticatorText",
			Placeholder:                         "Placeholder",
			RememberMeText:                      "RememberMeText",
			ResendActionText:                    "ResendActionText",
			ResendText:                          "ResendActionText",
			Title:                               "Title",
			InvalidOtpCodeFormat:                "InvalidOtpCodeFormat",
			InvalidCode:                         "InvalidCode",
			InvalidExpiredCode:                  "InvalidExpiredCode",
			AuthenticatorError:                  "AuthenticatorError",
			NoTransactionInProgress:             "NoTransactionInProgress",
			TooManyEmail:                        "TooManyEmail",
			TransactionNotFound:                 "TransactionNotFound",
			MfaEmailChallengeAuthenticatorError: "MfaEmailChallengeAuthenticatorError",
		},
		MfaEmailList: &ScreenMfaEmailList{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
	}

	err = m.Prompt.ReplacePromptMfaEmail(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfaEmail(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfaOtp(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfaOtp(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfaOtp(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfaOtp(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfaOtp{
		MfaOtpEnrollmentQr: &ScreenMfaOtpEnrollmentQr{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			ButtonText:            "ButtonText",
			CodeEnrollmentText:    "CodeEnrollmentText",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			//Documented but not supported
			//SeparatorText: "SeparatorText",
			InvalidOtpCodeFormat: "InvalidOtpCodeFormat",
			InvalidCode:          "InvalidCode",
			InvalidExpiredCode:   "InvalidExpiredCode",
			TooManyFailures:      "TooManyFailures",
			TransactionNotFound:  "TransactionNotFound",
			UserAlreadyEnrolled:  "UserAlreadyEnrolled",
		},
		MfaOtpEnrollmentCode: &ScreenMfaOtpEnrollmentCode{
			PageTitle:             "PageTitle",
			BackText:              "BackText",
			ButtonText:            "ButtonText",
			AltText:               "AltText",
			CopyCodeButtonText:    "CopyCodeButtonText",
			Description:           "Description",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			Title:                 "Title",
			TooManyFailures:       "TooManyFailures",
			TransactionNotFound:   "TransactionNotFound",
		},
		MfaOtpEnrollmentChallenge: &ScreenMfaOtpEnrollmentChallenge{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			ButtonText:            "ButtonText",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			RememberMeText:        "RememberMeText",
			AuthenticatorError:    "AuthenticatorError",
			TooManyFailures:       "TooManyFailures",
			TransactionNotFound:   "TransactionNotFound",
		},
	}

	err = m.Prompt.ReplacePromptMfaOtp(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfaOtp(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfaPhone(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfaPhone(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfaPhone(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfaPhone(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfaPhone{
		MfaPhoneChallenge: &ScreenMfaPhoneChallenge{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			ContinueButtonText:    "ContinueButtonText",
			SmsButtonText:         "SmsButtonText",
			VoiceButtonText:       "VoiceButtonText",
			ChooseMessageTypeText: "ChooseMessageTypeText",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			SendSmsFailed:         "SendSmsFailed",
			SendVoiceFailed:       "SendVoiceFailed",
			InvalidPhoneFormat:    "InvalidPhoneFormat",
			TooManySms:            "TooManySms",
			TooManyVoice:          "TooManyVoice",
			TransactionNotFound:   "TransactionNotFound",
			NoPhone:               "NoPhone",
		},
		MfaPhoneEnrollment: &ScreenMfaPhoneEnrollment{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			ContinueButtonText:    "ContinueButtonText",
			SmsButtonText:         "SmsButtonText",
			VoiceButtonText:       "VoiceButtonText",
			ChooseMessageTypeText: "ChooseMessageTypeText",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			SendSmsFailed:         "SendSmsFailed",
			SendVoiceFailed:       "SendVoiceFailed",
			InvalidPhoneFormat:    "InvalidPhoneFormat",
			TooManySms:            "TooManySms",
			TooManyVoice:          "TooManyVoice",
			TransactionNotFound:   "TransactionNotFound",
			NoPhone:               "NoPhone",
		},
	}

	err = m.Prompt.ReplacePromptMfaPhone(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfaPhone(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfaPush(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfaPush(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfaPush(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfaPush(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfaPush{
		MfaPushWelcome: &ScreenMfaPushWelcome{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			AndroidButtonText:     "AndroidButtonText",
			ButtonText:            "ButtonText",
			IosButtonText:         "IosButtonText",
			PickAuthenticatorText: "PickAuthenticatorText",
		},
		MfaPushEnrollmentQr: &ScreenMfaPushEnrollmentQr{
			PageTitle:                    "PageTitle",
			Title:                        "Title",
			Description:                  "Description",
			PickAuthenticatorText:        "PickAuthenticatorText",
			ButtonText:                   "ButtonText",
			EnrollmentTransactionPending: "EnrollmentTransactionPending",
		},
		MfaPushChallengePush: &ScreenMfaPushChallengePush{
			PageTitle:   "PageTitle",
			Title:       "Title",
			Description: "Description",
			//Documented but not supported
			//AwaitingConfirmation: "AwaitingConfirmation",
			ButtonText:                         "ButtonText",
			PickAuthenticatorText:              "PickAuthenticatorText",
			RememberMeText:                     "RememberMeText",
			ResendActionText:                   "ResendActionText",
			ResendText:                         "ResendText",
			EnterOtpCode:                       "EnterOtpCode",
			SeparatorText:                      "SeparatorText",
			ChallengeTransactionPending:        "ChallengeTransactionPending",
			PollingIntervalExceeded:            "PollingIntervalExceeded",
			TooManyPush:                        "TooManyPush",
			TransactionNotFound:                "TransactionNotFound",
			MfaPushVerifyTransactionPending:    "MfaPushVerifyTransactionPending",
			MfaPushVerifyAuthenticatorError:    "MfaPushVerifyAuthenticatorError",
			MfaPushChallengeAuthenticatorError: "MfaPushChallengeAuthenticatorError",
			TransactionRejected:                "TransactionRejected",
		},
		MfaPushList: &ScreenMfaPushList{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
	}

	err = m.Prompt.ReplacePromptMfaPush(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfaPush(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfaRecoveryCode(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfaRecoveryCode(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfaRecoveryCode(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfaRecoveryCode(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfaRecoveryCode{
		MfaRecoveryCodeEnrollment: &ScreenMfaRecoveryCodeEnrollment{
			PageTitle:          "PageTitle",
			Title:              "Title",
			Description:        "Description",
			AltText:            "AltText",
			ButtonText:         "ButtonText",
			CheckboxText:       "CheckboxText",
			CopyCodeButtonText: "CopyCodeButtonText",
			NoConfirmation:     "NoConfirmation",
		},
		MfaRecoveryCodeChallenge: &ScreenMfaRecoveryCodeChallenge{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			ButtonText:            "ButtonText",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			InvalidCode:           "InvalidCode",
			InvalidCodeFormat:     "InvalidCodeFormat",
			InvalidExpiredCode:    "InvalidExpiredCode",
			AuthenticatorError:    "AuthenticatorError",
			NoConfirmation:        "NoConfirmation",
			TooManyFailures:       "TooManyFailures",
			TransactionNotFound:   "TransactionNotFound",
		},
	}

	err = m.Prompt.ReplacePromptMfaRecoveryCode(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfaRecoveryCode(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfaSms(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfaSms(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfaSms(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfaSms(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfaSms{
		MfaCountryCodes: &ScreenMfaCountryCodes{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
		MfaSmsEnrollment: &ScreenMfaSmsEnrollment{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			ButtonText:            "ButtonText",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			SendSmsFailed:         "SendSmsFailed",
			InvalidPhoneFormat:    "InvalidPhoneFormat",
			InvalidPhone:          "InvalidPhone",
			TooManySms:            "TooManySms",
			TransactionNotFound:   "TransactionNotFound",
			NoPhone:               "NoPhone",
		},
		MfaSmsChallenge: &ScreenMfaSmsChallenge{
			PageTitle:                            "PageTitle",
			Title:                                "Title",
			Description:                          "Description",
			ButtonText:                           "ButtonText",
			EditText:                             "EditText",
			PickAuthenticatorText:                "PickAuthenticatorText",
			Placeholder:                          "Placeholder",
			RememberMeText:                       "RememberMeText",
			ResendActionText:                     "ResendActionText",
			ResendText:                           "ResendText",
			ResendVoiceActionSeparatorTextBefore: "ResendVoiceActionSeparatorTextBefore",
			ResendVoiceActionText:                "ResendVoiceActionText",
			ResendVoiceActionSeparatorTextAfter:  "ResendVoiceActionSeparatorTextAfter",
			InvalidOtpCodeFormat:                 "InvalidOtpCodeFormat",
			InvalidCode:                          "InvalidCode",
			InvalidExpiredCode:                   "InvalidExpiredCode",
			SendSmsFailed:                        "SendSmsFailed",
			AuthenticatorError:                   "AuthenticatorError",
			SmsAuthenticatorError:                "SmsAuthenticatorError",
			NoTransactionInProgress:              "NoTransactionInProgress",
			TooManyFailures:                      "TooManyFailures",
			TooManySms:                           "TooManySms",
			TransactionNotFound:                  "TransactionNotFound",
		},
		MfaSmsList: &ScreenMfaSmsList{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
	}

	err = m.Prompt.ReplacePromptMfaSms(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfaSms(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptMfaVoice(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptMfaVoice(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptMfaVoice(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptMfaVoice(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptMfaVoice{
		MfaVoiceEnrollment: &ScreenMfaVoiceEnrollment{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			ButtonText:            "ButtonText",
			PickAuthenticatorText: "PickAuthenticatorText",
			Placeholder:           "Placeholder",
			SendSmsFailed:         "SendSmsFailed",
			InvalidPhoneFormat:    "InvalidPhoneFormat",
			InvalidPhone:          "InvalidPhone",
			TooManySms:            "TooManySms",
			TransactionNotFound:   "TransactionNotFound",
			NoPhone:               "NoPhone",
		},
		MfaVoiceChallenge: &ScreenMfaVoiceChallenge{
			PageTitle:                          "PageTitle",
			Title:                              "Title",
			Description:                        "Description",
			ButtonText:                         "ButtonText",
			EditText:                           "EditText",
			PickAuthenticatorText:              "PickAuthenticatorText",
			Placeholder:                        "Placeholder",
			RememberMeText:                     "RememberMeText",
			ResendActionText:                   "ResendActionText",
			ResendText:                         "ResendText",
			ResendSmsActionSeparatorTextBefore: "ResendSmsActionSeparatorTextBefore",
			ResendSmsActionText:                "ResendSmsActionText",
			ResendSmsActionSeparatorTextAfter:  "ResendSmsActionSeparatorTextAfter",
			InvalidOtpCodeFormat:               "InvalidOtpCodeFormat",
			InvalidCode:                        "InvalidCode",
			InvalidExpiredCode:                 "InvalidExpiredCode",
			SendVoiceFailed:                    "SendVoiceFailed",
			AuthenticatorError:                 "AuthenticatorError",
			VoiceAuthenticatorError:            "VoiceAuthenticatorError",
			NoTransactionInProgress:            "NoTransactionInProgress",
			TooManyFailures:                    "TooManyFailures",
			TooManyVoice:                       "TooManyVoice",
			TransactionNotFound:                "TransactionNotFound",
		},
	}

	err = m.Prompt.ReplacePromptMfaVoice(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptMfaVoice(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptOrganizations(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptOrganizations(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptOrganizations(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptOrganizations(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptOrganizationSelection{
		OrganizationSelection: &ScreenOrganizationSelection{
			PageTitle:   "PageTitle",
			Title:       "Title",
			Description: "Description",
			ButtonText:  "ButtonText",
			Placeholder: "Placeholder",
			//Documented but not supported
			//ErrorInvalidOrganization: "ErrorInvalidOrganization",
			LogoAltText: "LogoAltText",
		},
	}

	err = m.Prompt.ReplacePromptOrganizations(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptOrganizations(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptResetPassword(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptResetPassword(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptResetPassword(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptResetPassword(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptResetPassword{
		ResetPasswordRequest: &ScreenResetPasswordRequest{
			PageTitle:               "PageTitle",
			Title:                   "Title",
			BackToLoginLinkText:     "BackToLoginLinkText",
			ButtonText:              "ButtonText",
			DescriptionEmail:        "DescriptionEmail",
			DescriptionUsername:     "DescriptionUsername",
			PlaceholderEmail:        "PlaceholderEmail",
			PlaceholderUsername:     "PlaceholderUsername",
			InvalidEmailFormat:      "InvalidEmailFormat",
			Auth0UsersExpiredTicket: "Auth0UsersExpiredTicket",
			CustomScriptErrorCode:   "CustomScriptErrorCode",
			Auth0UsersUsedTicket:    "Auth0UsersUsedTicket",
			ResetPasswordError:      "ResetPasswordError",
			TooManyEmail:            "TooManyEmail",
			TooManyRequests:         "TooManyRequests",
			NoEmail:                 "NoEmail",
			NoUsername:              "NoUsername",
		},
		ResetPasswordEmail: &ScreenResetPasswordEmail{
			PageTitle:        "PageTitle",
			Title:            "Title",
			EmailDescription: "EmailDescription",
			ResendLinkText:   "ResendLinkText",
			//Documented but not supported
			//ResendText:              "ResendText",
			UsernameDescription: "UsernameDescription",
		},
		ResetPassword: &ScreenResetPassword{
			PageTitle:                  "PageTitle",
			Title:                      "Title",
			Description:                "Description",
			ButtonText:                 "ButtonText",
			PasswordPlaceholder:        "PasswordPlaceholder",
			ReEnterPasswordPlaceholder: "ReEnterPasswordPlaceholder",
			PasswordSecurityText:       "PasswordSecurityText",
			Auth0UsersExpiredTicket:    "Auth0UsersExpiredTicket",
			CustomScriptErrorCode:      "CustomScriptErrorCode",
			Auth0UsersUsedTicket:       "Auth0UsersUsedTicket",
			Auth0UsersValidation:       "Auth0UsersValidation",
			NoReEnterPassword:          "NoReEnterPassword",
		},
		ResetPasswordSuccess: &ScreenResetPasswordSuccess{
			PageTitle:   "PageTitle",
			EventTitle:  "EventTitle",
			Description: "Description",
			ButtonText:  "ButtonText",
		},
		ResetPasswordError: &ScreenResetPasswordError{
			PageTitle:               "PageTitle",
			BackToLoginLinkText:     "BackToLoginLinkText",
			DescriptionExpired:      "DescriptionExpired",
			DescriptionGeneric:      "DescriptionGeneric",
			DescriptionUsed:         "DescriptionUsed",
			EventTitleExpired:       "EventTitleExpired",
			EventTitleUsed:          "EventTitleUsed",
			Auth0UsersExpiredTicket: "Auth0UsersExpiredTicket",
			CustomScriptErrorCode:   "CustomScriptErrorCode",
			Auth0UsersUsedTicket:    "Auth0UsersUsedTicket",
			Auth0UsersValidation:    "Auth0UsersValidation",
			ResetPasswordError:      "ResetPasswordError",
		},
	}

	err = m.Prompt.ReplacePromptResetPassword(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptResetPassword(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptSignup(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptSignup(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptSignup(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptSignup(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptSignup{
		Signup: &ScreenSignup{
			PageTitle:                        "PageTitle",
			Title:                            "Title",
			Description:                      "Description",
			SeparatorText:                    "SeparatorText",
			ButtonText:                       "ButtonText",
			EmailPlaceholder:                 "EmailPlaceholder",
			FederatedConnectionButtonText:    "FederatedConnectionButtonText",
			LoginActionLinkText:              "LoginActionLinkText",
			LoginActionText:                  "LoginActionText",
			PasswordPlaceholder:              "PasswordPlaceholder",
			UsernamePlaceholder:              "UsernamePlaceholder",
			EmailInUse:                       "EmailInUse",
			InvalidEmailFormat:               "InvalidEmailFormat",
			PasswordTooWeak:                  "PasswordTooWeak",
			PasswordTooCommon:                "PasswordTooCommon",
			PasswordPreviouslyUsed:           "PasswordPreviouslyUsed",
			PasswordMismatch:                 "PasswordMismatch",
			InvalidUsername:                  "InvalidUsername",
			InvalidUsernameMaxLength:         "InvalidUsernameMaxLength",
			InvalidUsernameMinLength:         "InvalidUsernameMinLength",
			InvalidUsernameInvalidCharacters: "InvalidUsernameInvalidCharacters",
			InvalidUsernameEmailNotAllowed:   "InvalidUsernameEmailNotAllowed",
			UsernameTaken:                    "UsernameTaken",
			CustomScriptErrorCode:            "CustomScriptErrorCode",
			Auth0UsersValidation:             "Auth0UsersValidation",
			InvalidConnection:                "InvalidConnection",
			IpBlocked:                        "IpBlocked",
			IpSignupBlocked:                  "IpSignupBlocked",
			NoDbConnection:                   "NoDbConnection",
			NoEmail:                          "NoEmail",
			NoPassword:                       "NoPassword",
			NoReEnterPassword:                "NoReEnterPassword",
			NoUsername:                       "NoUsername",
		},
	}

	err = m.Prompt.ReplacePromptSignup(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptSignup(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptSignupId(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptSignupId(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptSignupId(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptSignupId(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptSignupId{
		SignupId: &ScreenSignupId{
			PageTitle:                        "PageTitle",
			Title:                            "Title",
			Description:                      "Description",
			SeparatorText:                    "SeparatorText",
			ButtonText:                       "ButtonText",
			EmailPlaceholder:                 "EmailPlaceholder",
			FederatedConnectionButtonText:    "FederatedConnectionButtonText",
			LoginActionLinkText:              "LoginActionLinkText",
			LoginActionText:                  "LoginActionText",
			PasswordPlaceholder:              "PasswordPlaceholder",
			PasswordSecurityText:             "PasswordSecurityText",
			UsernamePlaceholder:              "UsernamePlaceholder",
			LogoAltText:                      "LogoAltText",
			EmailInUse:                       "EmailInUse",
			InvalidEmailFormat:               "InvalidEmailFormat",
			PasswordTooWeak:                  "PasswordTooWeak",
			PasswordTooCommon:                "PasswordTooCommon",
			PasswordPreviouslyUsed:           "PasswordPreviouslyUsed",
			PasswordMismatch:                 "PasswordMismatch",
			InvalidUsername:                  "InvalidUsername",
			InvalidUsernameMaxLength:         "InvalidUsernameMaxLength",
			InvalidUsernameMinLength:         "InvalidUsernameMinLength",
			InvalidUsernameInvalidCharacters: "InvalidUsernameInvalidCharacters",
			InvalidUsernameEmailNotAllowed:   "InvalidUsernameEmailNotAllowed",
			UsernameTaken:                    "UsernameTaken",
			CustomScriptErrorCode:            "CustomScriptErrorCode",
			Auth0UsersValidation:             "Auth0UsersValidation",
			InvalidConnection:                "InvalidConnection",
			IpBlocked:                        "IpBlocked",
			IpSignupBlocked:                  "IpSignupBlocked",
			NoDbConnection:                   "NoDbConnection",
			NoEmail:                          "NoEmail",
			NoPassword:                       "NoPassword",
			NoReEnterPassword:                "NoReEnterPassword",
			NoUsername:                       "NoUsername",
		},
	}

	err = m.Prompt.ReplacePromptSignupId(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptSignupId(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestPromptSignupPassword(t *testing.T) {
	StartVal, err := m.Prompt.ReadPromptSignupPassword(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.Prompt.ReplacePromptSignupPassword(Language, StartVal)
		if err != nil {
			t.Fatal(err)
		}
		ReturnVal, err := m.Prompt.ReadPromptSignupPassword(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptSignupPassword{
		SignupPassword: &ScreenSignupPassword{
			PageTitle:                     "PageTitle",
			Title:                         "Title",
			Description:                   "Description",
			SeparatorText:                 "SeparatorText",
			ButtonText:                    "ButtonText",
			EmailPlaceholder:              "EmailPlaceholder",
			FederatedConnectionButtonText: "FederatedConnectionButtonText",
			LoginActionLinkText:           "LoginActionLinkText",
			LoginActionText:               "LoginActionText",
			//Documented but not supported
			//LoginLinkText: "LoginLinkText",
			InvitationTitle:                  "InvitationTitle",
			InvitationDescription:            "InvitationDescription",
			PasswordPlaceholder:              "PasswordPlaceholder",
			PasswordSecurityText:             "PasswordSecurityText",
			UsernamePlaceholder:              "UsernamePlaceholder",
			EmailInUse:                       "EmailInUse",
			InvalidEmailFormat:               "InvalidEmailFormat",
			PasswordTooWeak:                  "PasswordTooWeak",
			PasswordTooCommon:                "PasswordTooCommon",
			PasswordPreviouslyUsed:           "PasswordPreviouslyUsed",
			PasswordMismatch:                 "PasswordMismatch",
			InvalidUsername:                  "InvalidUsername",
			InvalidUsernameMaxLength:         "InvalidUsernameMaxLength",
			InvalidUsernameMinLength:         "InvalidUsernameMinLength",
			InvalidUsernameInvalidCharacters: "InvalidUsernameInvalidCharacters",
			InvalidUsernameEmailNotAllowed:   "InvalidUsernameEmailNotAllowed",
			UsernameTaken:                    "UsernameTaken",
			CustomScriptErrorCode:            "CustomScriptErrorCode",
			Auth0UsersValidation:             "Auth0UsersValidation",
			InvalidConnection:                "InvalidConnection",
			IpBlocked:                        "IpBlocked",
			IpSignupBlocked:                  "IpSignupBlocked",
			NoDbConnection:                   "NoDbConnection",
			NoEmail:                          "NoEmail",
			NoPassword:                       "NoPassword",
			NoReEnterPassword:                "NoReEnterPassword",
			NoUsername:                       "NoUsername",
			LogoAltText:                      "LogoAltText",
		},
	}

	err = m.Prompt.ReplacePromptSignupPassword(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.Prompt.ReadPromptSignupPassword(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}
