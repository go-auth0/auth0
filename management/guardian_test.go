package management

import (
	"testing"

	"gopkg.in/auth0.v5"
)

func TestGuardian(t *testing.T) {

	t.Run("MultiFactor", func(t *testing.T) {

		t.Run("List", func(t *testing.T) {
			mfa, err := m.Guardian.MultiFactor.List()
			if err != nil {
				t.Error(err)
			}
			t.Logf("%v\n", mfa)
		})

		t.Run("SMS", func(t *testing.T) {

			t.Run("Enable", func(t *testing.T) {
				defer m.Guardian.MultiFactor.SMS.Enable(false)

				err := m.Guardian.MultiFactor.SMS.Enable(true)
				if err != nil {
					t.Error(err)
				}

				mfa, _ := m.Guardian.MultiFactor.List()
				t.Logf("%v\n", mfa)
			})

			t.Run("Template", func(t *testing.T) {
				defer m.Guardian.MultiFactor.SMS.UpdateTemplate(&MultiFactorSMSTemplate{
					EnrollmentMessage:   auth0.String(""),
					VerificationMessage: auth0.String(""),
				})

				err := m.Guardian.MultiFactor.SMS.UpdateTemplate(&MultiFactorSMSTemplate{
					EnrollmentMessage:   auth0.String("Test {{code}} for {{tenant.friendly_name}}"),
					VerificationMessage: auth0.String("Test {{code}} for {{tenant.friendly_name}}"),
				})
				if err != nil {
					t.Error(err)
				}

				template, err := m.Guardian.MultiFactor.SMS.Template()
				if err != nil {
					t.Error(err)
				}
				t.Logf("%v\n", template)
			})

			t.Run("Twilio", func(t *testing.T) {
				defer m.Guardian.MultiFactor.SMS.UpdateTwilio(&MultiFactorProviderTwilio{
					From:      auth0.String(""),
					AuthToken: auth0.String(""),
					SID:       auth0.String(""),
				})

				err := m.Guardian.MultiFactor.SMS.UpdateTwilio(&MultiFactorProviderTwilio{
					From:      auth0.String("123456789"),
					AuthToken: auth0.String("test_token"),
					SID:       auth0.String("test_sid"),
				})
				if err != nil {
					t.Error(err)
				}

				twilio, err := m.Guardian.MultiFactor.SMS.Twilio()
				if err != nil {
					t.Error(err)
				}
				t.Logf("%v\n", twilio)
			})
		})

		t.Run("Push", func(t *testing.T) {

			t.Run("Enable", func(t *testing.T) {
				defer m.Guardian.MultiFactor.Push.Enable(false)

				err := m.Guardian.MultiFactor.Push.Enable(true)
				if err != nil {
					t.Error(err)
				}

				mfa, _ := m.Guardian.MultiFactor.List()
				t.Logf("%v\n", mfa)
			})

			t.Run("AmazonSNS", func(t *testing.T) {
				defer m.Guardian.MultiFactor.Push.UpdateAmazonSNS(&MultiFactorProviderAmazonSNS{
					AccessKeyID:                auth0.String(""),
					SecretAccessKeyID:          auth0.String(""),
					Region:                     auth0.String(""),
					APNSPlatformApplicationARN: auth0.String(""),
					GCMPlatformApplicationARN:  auth0.String(""),
				})

				err := m.Guardian.MultiFactor.Push.UpdateAmazonSNS(&MultiFactorProviderAmazonSNS{
					AccessKeyID:                auth0.String("test"),
					SecretAccessKeyID:          auth0.String("test_secret"),
					Region:                     auth0.String("us-west-1"),
					APNSPlatformApplicationARN: auth0.String("test_arn"),
					GCMPlatformApplicationARN:  auth0.String("test_arn"),
				})
				if err != nil {
					t.Error(err)
				}

				sns, err := m.Guardian.MultiFactor.Push.AmazonSNS()
				if err != nil {
					t.Error(err)
				}
				t.Logf("%v\n", sns)
			})
		})

		t.Run("Email", func(t *testing.T) {

			t.Run("Enable", func(t *testing.T) {
				defer m.Guardian.MultiFactor.Email.Enable(false)

				err := m.Guardian.MultiFactor.Email.Enable(true)
				if err != nil {
					t.Error(err)
				}

				mfa, _ := m.Guardian.MultiFactor.List()
				t.Logf("%v\n", mfa)
			})
		})

		t.Run("DUO", func(t *testing.T) {

			t.Run("Enable", func(t *testing.T) {
				defer m.Guardian.MultiFactor.DUO.Enable(false)

				err := m.Guardian.MultiFactor.DUO.Enable(true)
				if err != nil {
					t.Error(err)
				}

				mfa, _ := m.Guardian.MultiFactor.List()
				t.Logf("%v\n", mfa)
			})
		})

		t.Run("OTP", func(t *testing.T) {

			t.Run("Enable", func(t *testing.T) {
				defer m.Guardian.MultiFactor.OTP.Enable(false)

				err := m.Guardian.MultiFactor.OTP.Enable(true)
				if err != nil {
					t.Error(err)
				}

				mfa, _ := m.Guardian.MultiFactor.List()
				t.Logf("%v\n", mfa)
			})
		})
	})
}
