package management

import (
	"testing"

	"gopkg.in/auth0.v1"
)

func TestGuardian(t *testing.T) {
	testFactor := OtpGuardianFactor

	defer func() {
		emptyString := ""
		m.Guardian.UpdateFactor(testFactor, &GuardianFactor{
			Enabled: auth0.Bool(false),
		})
		m.Guardian.UpdateSmsTemplate(&GuardianSmsTemplate{
			EnrollmentMessage:   auth0.String("{{code}} is your verification code for {{tenant.friendly_name}}. Please enter this code to verify your enrollment."),
			VerificationMessage: auth0.String("{{code}} is your verification code for {{tenant.friendly_name}}"),
		})
		m.Guardian.UpdatePushNotificationSnsConfig(&GuardianPushNotificationSnsConfig{
			AwsAccessKeyID:                &emptyString,
			AwsSecretAccessKeyID:          &emptyString,
			AwsRegion:                     &emptyString,
			SnsApnsPlatformApplicationArn: &emptyString,
			SnsGcmPlatformApplicationArn:  &emptyString,
		})
		m.Guardian.UpdateSmsTwilioConfig(&GuardianSmsTwilioConfig{
			From:                &emptyString,
			MessagingServiceSid: &emptyString,
			AuthToken:           &emptyString,
			Sid:                 &emptyString,
		})
	}()

	t.Run("ListFactors", func(t *testing.T) {
		gf, err := m.Guardian.ListFactors()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", gf)
	})

	t.Run("UpdateFactor", func(t *testing.T) {
		err := m.Guardian.UpdateFactor(testFactor, &GuardianFactor{
			Enabled: auth0.Bool(true),
		})
		if err != nil {
			t.Error(err)
		}
		gf, _ := m.Guardian.ListFactors()
		t.Logf("%v\n", gf)
	})

	t.Run("GetSmsTemplate", func(t *testing.T) {
		st, err := m.Guardian.GetSmsTemplate()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", st)
	})

	t.Run("UpdateSmsTemplate", func(t *testing.T) {
		err := m.Guardian.UpdateSmsTemplate(&GuardianSmsTemplate{
			EnrollmentMessage:   auth0.String("Test {{code}} for {{tenant.friendly_name}}"),
			VerificationMessage: auth0.String("Test {{code}} for {{tenant.friendly_name}}"),
		})
		if err != nil {
			t.Error(err)
		}
		st, _ := m.Guardian.GetSmsTemplate()
		t.Logf("%v\n", st)
	})

	t.Run("GetPushNotificationSnsConfig", func(t *testing.T) {
		sc, err := m.Guardian.GetPushNotificationSnsConfig()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", sc)
	})

	t.Run("UpdatePushNotificationSnsConfig", func(t *testing.T) {
		err := m.Guardian.UpdatePushNotificationSnsConfig(&GuardianPushNotificationSnsConfig{
			AwsAccessKeyID:                auth0.String("test"),
			AwsSecretAccessKeyID:          auth0.String("test_secret"),
			AwsRegion:                     auth0.String("us-west-1"),
			SnsApnsPlatformApplicationArn: auth0.String("test_arn"),
			SnsGcmPlatformApplicationArn:  auth0.String("test_arn"),
		})
		if err != nil {
			t.Error(err)
		}
		sc, _ := m.Guardian.GetPushNotificationSnsConfig()
		t.Logf("%v\n", sc)
	})

	t.Run("GetSmsTwilioConfig", func(t *testing.T) {
		tc, err := m.Guardian.GetSmsTwilioConfig()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", tc)
	})

	t.Run("UpdateSmsTwilioConfig", func(t *testing.T) {
		err := m.Guardian.UpdateSmsTwilioConfig(&GuardianSmsTwilioConfig{
			From:      auth0.String("123456789"),
			AuthToken: auth0.String("test_token"),
			Sid:       auth0.String("test_sid"),
		})
		if err != nil {
			t.Error(err)
		}
		tc, _ := m.Guardian.GetSmsTwilioConfig()
		t.Logf("%v\n", tc)
	})
}
