package management

import (
	"github.com/yieldr/go-auth0"
	"testing"
)

func TestJob(t *testing.T) {

	var err error

	u := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("example@example.com"),
		Password:   auth0.String("I have a password and its a secret"),
	}
	err = m.User.Create(u)
	if err != nil {
		t.Error(err)
	}

	userID := auth0.StringValue(u.ID)

	defer m.User.Delete(userID)

	t.Run("verification email", func(t *testing.T) {
		job := &Job{
			UserID: auth0.String(userID),
		}

		err = m.Job.VerifyEmail(
			job,
		)
		if err != nil {
			t.Error(err)
		}

		if auth0.StringValue(job.Type) != "verification_email" {
			t.Errorf("expected job type to be verification_email, got %s", auth0.StringValue(job.Type))
		}
	})
}
