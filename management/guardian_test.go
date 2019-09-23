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
		m.Guardian.UpdateFactorPushNotificationSnsConfig(&GuardianFactorPushNotificationSnsConfig{
			AwsAccessKeyID:                &emptyString,
			AwsSecretAccessKeyID:          &emptyString,
			AwsRegion:                     &emptyString,
			SnsApnsPlatformApplicationArn: &emptyString,
			SnsGcmPlatformApplicationArn:  &emptyString,
		})
		m.Guardian.UpdateFactorSmsTwilioConfig(&GuardianFactorSmsTwilioConfig{
			From:                &emptyString,
			MessagingServiceSid: &emptyString,
			AuthToken:           &emptyString,
			Sid:                 &emptyString,
		})
	}()

	t.Run("ListFactorsAndStatuses", func(t *testing.T) {
		gf, err := m.Guardian.ListFactorsAndStatuses()
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
		gf, _ := m.Guardian.ListFactorsAndStatuses()
		t.Logf("%v\n", gf)
	})

	t.Run("GetFactorPushNotificationSnsConfig", func(t *testing.T) {
		sc, err := m.Guardian.GetFactorPushNotificationSnsConfig()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", sc)
	})

	t.Run("UpdateFactorPushNotificationSnsConfig", func(t *testing.T) {
		err := m.Guardian.UpdateFactorPushNotificationSnsConfig(&GuardianFactorPushNotificationSnsConfig{
			AwsAccessKeyID:                auth0.String("test"),
			AwsSecretAccessKeyID:          auth0.String("test_secret"),
			AwsRegion:                     auth0.String("us-west-1"),
			SnsApnsPlatformApplicationArn: auth0.String("test_arn"),
			SnsGcmPlatformApplicationArn:  auth0.String("test_arn"),
		})
		if err != nil {
			t.Error(err)
		}
		sc, _ := m.Guardian.GetFactorPushNotificationSnsConfig()
		t.Logf("%v\n", sc)
	})

	t.Run("GetFactorSmsTwilioConfig", func(t *testing.T) {
		tc, err := m.Guardian.GetFactorSmsTwilioConfig()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", tc)
	})

	t.Run("UpdateFactorSmsTwilioConfig", func(t *testing.T) {
		err := m.Guardian.UpdateFactorSmsTwilioConfig(&GuardianFactorSmsTwilioConfig{
			From:      auth0.String("123456789"),
			AuthToken: auth0.String("test_token"),
			Sid:       auth0.String("test_sid"),
		})
		if err != nil {
			t.Error(err)
		}
		tc, _ := m.Guardian.GetFactorSmsTwilioConfig()
		t.Logf("%v\n", tc)
	})
}
