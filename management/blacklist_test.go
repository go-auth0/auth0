package management

import (
	"testing"
)

func TestPrompt(t *testing.T) {
	t.Run("GetSettings", func(t *testing.T) {
		bl, err := m.Blacklist.GetBlacklistedTokens()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", bl)
	})

	t.Run("UpdateSettings", func(t *testing.T) {
		err := m.Blacklist.BlacklistToken(&BlacklistToken{
			Aud: "test",
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
