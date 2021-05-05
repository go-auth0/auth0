package management

import (
	"fmt"
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

// CustomText tests
//
// instructions:
// install go: https://golang.org/doc/install
// $ cd auth0/management
// $ go test -run CustomText

// Language to use with tests
const Language = "en"

func TestConsentCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadConsent(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetConsent(Language, StartVal)
		if err != nil {
			// fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		}
		ReturnVal, err := m.CustomText.ReadConsent(Language)
		if err != nil {
			t.Fatal(err)
		}
		expect.Expect(t, ReturnVal, StartVal)
	})

	SetVal := PromptConsent{
		Consent: ScreenConsent{
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

	err = m.CustomText.SetConsent(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadConsent(Language)
	if err != nil {
		t.Fatal(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestDeviceFlowCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadDeviceFlow(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetDeviceFlow(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadDeviceFlow(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptDeviceFlow{
		DeviceCodeActivation: ScreenDeviceCodeActivation{
			PageTitle:          "PageTitle",
			ButtonText:         "ButtonText",
			Description:        "Description",
			Placeholder:        "Placeholder",
			Title:              "Title",
			InvalidExpiredCode: "InvalidExpiredCode",
			NoCode:             "NoCode",
			InvalidCode:        "InvalidCode",
		},
		DeviceCodeActivationAllowed: ScreenDeviceCodeActivationAllowed{
			PageTitle:   "PageTitle",
			Description: "Description",
			EventTitle:  "EventTitle",
		},
		DeviceCodeActivationDenied: ScreenDeviceCodeActivationDenied{
			PageTitle:   "PageTitle",
			Description: "Description",
			EventTitle:  "EventTitle",
		},
		DeviceCodeConfirmation: ScreenDeviceCodeConfirmation{
			PageTitle:         "PageTitle",
			Description:       "Description",
			InputCodeLabel:    "InputCodeLabel",
			Title:             "Title",
			ConfirmButtonText: "ConfirmButtonText",
			CancelButtonText:  "CancelButtonText",
			ConfirmationText:  "ConfirmationText",
		},
	}

	err = m.CustomText.SetDeviceFlow(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadDeviceFlow(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestEmailOtpChallengeCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadEmailOtpChallenge(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetEmailOtpChallenge(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadEmailOtpChallenge(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptEmailOtpChallenge{
		EmailOtpChallenge: ScreenEmailOtpChallenge{
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

	err = m.CustomText.SetEmailOtpChallenge(Language, &SetVal)
	if err != nil {
		t.Fatal(err)
	}

	ReturnVal, err := m.CustomText.ReadEmailOtpChallenge(Language)
	if err != nil {
		t.Fatal(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestEmailVerificationCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadEmailVerification(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetEmailVerification(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadEmailVerification(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptEmailVerification{
		EmailVerificationResult: ScreenEmailVerificationResult{
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

	err = m.CustomText.SetEmailVerification(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadEmailVerification(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestInvitationCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadInvitation(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetInvitation(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadInvitation(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptInvitation{
		AcceptInvitation: ScreenAcceptInvitation{
			PageTitle:   "PageTitle",
			Title:       "Title",
			Description: "Description",
			ButtonText:  "ButtonText",
			LogoAltText: "LogoAltText",
		},
	}

	err = m.CustomText.SetInvitation(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadInvitation(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestLoginCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadLogin(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetLogin(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadLogin(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptLogin{
		Login: ScreenLogin{
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

	err = m.CustomText.SetLogin(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadLogin(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestLoginIdCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadLoginId(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetLoginId(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadLoginId(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptLoginId{
		LoginId: ScreenLoginId{
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

	err = m.CustomText.SetLoginId(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadLoginId(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestLoginPasswordCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadLoginPassword(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetLoginPassword(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadLoginPassword(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptLoginPassword{
		LoginPassword: ScreenLoginPassword{
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

	err = m.CustomText.SetLoginPassword(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadLoginPassword(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestLoginEmailVerificationCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadLoginEmailVerification(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetLoginEmailVerification(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadLoginEmailVerification(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptLoginEmailVerification{
		LoginEmailVerification: ScreenLoginEmailVerification{
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

	err = m.CustomText.SetLoginEmailVerification(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadLoginEmailVerification(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfa(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfa(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfa(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfa{
		MfaEnrollResult: ScreenMfaEnrollResult{
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
		MfaLoginOptions: ScreenMfaLoginOptions{
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
		MfaBeginEnrollOptions: ScreenMfaBeginEnrollOptions{
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

	err = m.CustomText.SetMfa(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfa(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaEmailCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfaEmail(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfaEmail(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfaEmail(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfaEmail{
		MfaEmailChallenge: ScreenMfaEmailChallenge{
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
		MfaEmailList: ScreenMfaEmailList{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
	}

	err = m.CustomText.SetMfaEmail(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfaEmail(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaOtpCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfaOtp(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfaOtp(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfaOtp(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfaOtp{
		MfaOtpEnrollmentQr: ScreenMfaOtpEnrollmentQr{
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
		MfaOtpEnrollmentCode: ScreenMfaOtpEnrollmentCode{
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
		MfaOtpEnrollmentChallenge: ScreenMfaOtpEnrollmentChallenge{
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

	err = m.CustomText.SetMfaOtp(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfaOtp(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaPhoneCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfaPhone(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfaPhone(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfaPhone(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfaPhone{
		MfaPhoneChallenge: ScreenMfaPhoneChallenge{
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
		MfaPhoneEnrollment: ScreenMfaPhoneEnrollment{
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

	err = m.CustomText.SetMfaPhone(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfaPhone(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaPushCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfaPush(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfaPush(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfaPush(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfaPush{
		MfaPushWelcome: ScreenMfaPushWelcome{
			PageTitle:             "PageTitle",
			Title:                 "Title",
			Description:           "Description",
			AndroidButtonText:     "AndroidButtonText",
			ButtonText:            "ButtonText",
			IosButtonText:         "IosButtonText",
			PickAuthenticatorText: "PickAuthenticatorText",
		},
		MfaPushEnrollmentQr: ScreenMfaPushEnrollmentQr{
			PageTitle:                    "PageTitle",
			Title:                        "Title",
			Description:                  "Description",
			PickAuthenticatorText:        "PickAuthenticatorText",
			ButtonText:                   "ButtonText",
			EnrollmentTransactionPending: "EnrollmentTransactionPending",
		},
		MfaPushChallengePush: ScreenMfaPushChallengePush{
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
		MfaPushList: ScreenMfaPushList{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
	}

	err = m.CustomText.SetMfaPush(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfaPush(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaRecoveryCodeCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfaRecoveryCode(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfaRecoveryCode(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfaRecoveryCode(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfaRecoveryCode{
		MfaRecoveryCodeEnrollment: ScreenMfaRecoveryCodeEnrollment{
			PageTitle:          "PageTitle",
			Title:              "Title",
			Description:        "Description",
			AltText:            "AltText",
			ButtonText:         "ButtonText",
			CheckboxText:       "CheckboxText",
			CopyCodeButtonText: "CopyCodeButtonText",
			NoConfirmation:     "NoConfirmation",
		},
		MfaRecoveryCodeChallenge: ScreenMfaRecoveryCodeChallenge{
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

	err = m.CustomText.SetMfaRecoveryCode(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfaRecoveryCode(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaSmsCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfaSms(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfaSms(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfaSms(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfaSms{
		MfaCountryCodes: ScreenMfaCountryCodes{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
		MfaSmsEnrollment: ScreenMfaSmsEnrollment{
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
		MfaSmsChallenge: ScreenMfaSmsChallenge{
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
		MfaSmsList: ScreenMfaSmsList{
			PageTitle: "PageTitle",
			BackText:  "BackText",
			Title:     "Title",
		},
	}

	err = m.CustomText.SetMfaSms(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfaSms(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestMfaVoiceCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadMfaVoice(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetMfaVoice(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadMfaVoice(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptMfaVoice{
		MfaVoiceEnrollment: ScreenMfaVoiceEnrollment{
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
		MfaVoiceChallenge: ScreenMfaVoiceChallenge{
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

	err = m.CustomText.SetMfaVoice(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadMfaVoice(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestOrganizationsCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadOrganizations(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetOrganizations(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadOrganizations(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptOrganizations{
		OrganizationSelection: ScreenOrganizationSelection{
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

	err = m.CustomText.SetOrganizations(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadOrganizations(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestResetPasswordCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadResetPassword(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetResetPassword(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadResetPassword(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptResetPassword{
		ResetPasswordRequest: ScreenResetPasswordRequest{
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
		ResetPasswordEmail: ScreenResetPasswordEmail{
			PageTitle:        "PageTitle",
			Title:            "Title",
			EmailDescription: "EmailDescription",
			ResendLinkText:   "ResendLinkText",
			//Documented but not supported
			//ResendText:              "ResendText",
			UsernameDescription: "UsernameDescription",
		},
		ResetPassword: ScreenResetPassword{
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
		ResetPasswordSuccess: ScreenResetPasswordSuccess{
			PageTitle:   "PageTitle",
			EventTitle:  "EventTitle",
			Description: "Description",
			ButtonText:  "ButtonText",
		},
		ResetPasswordError: ScreenResetPasswordError{
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

	err = m.CustomText.SetResetPassword(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadResetPassword(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestSignupCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadSignup(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetSignup(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadSignup(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptSignup{
		Signup: ScreenSignup{
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
			NoPasswordd:                      "NoPasswordd",
			NoReEnterPassword:                "NoReEnterPassword",
			NoUsername:                       "NoUsername",
		},
	}

	err = m.CustomText.SetSignup(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadSignup(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestSignupIdCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadSignupId(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetSignupId(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadSignupId(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptSignupId{
		SignupId: ScreenSignupId{
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

	err = m.CustomText.SetSignupId(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadSignupId(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}

func TestSignupPasswordCustomText(t *testing.T) {
	StartVal, err := m.CustomText.ReadSignupPassword(Language)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = m.CustomText.SetSignupPassword(Language, StartVal)
		if err != nil {
			// cleanup fails if StartVal is empty object
			fmt.Printf("Warning, Set-StartVal failed: %q", err)
		} else {
			ReturnVal, err := m.CustomText.ReadSignupPassword(Language)
			if err != nil {
				t.Fatal(err)
			}
			expect.Expect(t, ReturnVal, StartVal)
		}
	})

	SetVal := PromptSignupPassword{
		SignupPassword: ScreenSignupPassword{
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

	err = m.CustomText.SetSignupPassword(Language, &SetVal)
	if err != nil {
		t.Error(err)
	}

	ReturnVal, err := m.CustomText.ReadSignupPassword(Language)
	if err != nil {
		t.Error(err)
	}
	expect.Expect(t, *ReturnVal, SetVal)
}
