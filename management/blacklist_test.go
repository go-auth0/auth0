package management

import (
	"testing"
	"time"

	"gopkg.in/auth0.v1"
)

func TestPrompt(t *testing.T) {
	c := &Client{
		Name: auth0.Stringf("Test Client - Client Grant (%s)", time.Now().Format(time.StampMilli)),
	}
	err := m.Client.Create(c)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Client.Delete(auth0.StringValue(c.ClientID))

	t.Run("GetSettings", func(t *testing.T) {
		bl, err := m.Blacklist.GetBlacklistedTokens()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", bl)
	})

	t.Run("UpdateSettings", func(t *testing.T) {
		err := m.Blacklist.BlacklistToken(&BlacklistToken{
			Aud: auth0.StringValue(c.ClientID),
			Jti: "test",
		})
		if err != nil {
			t.Error(err)
		}
		bl, err := m.Blacklist.GetBlacklistedTokens()
		if err != nil {
			t.Error(err)
		}
		if len(bl) == 0 {
			t.Error("unexpected output; blacklist should not be empty")
		}
		t.Logf("%v\n", bl)
	})
}
