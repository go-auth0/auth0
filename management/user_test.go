package management

import "testing"

func TestUser(t *testing.T) {

	u := &User{
		Connection: "Username-Password-Authentication",
		Email:      "chuck@chucknorris.com",
		Password:   "Passwords hide their Chuck",
		UserMetadata: map[string]interface{}{
			"favourite_attack": "roundhouse_kick",
		},
		EmailVerified: true,
		VerifyEmail:   false,
		AppMetadata: map[string]interface{}{
			"facts": []string{
				"count_to_infinity_twice",
				"kill_two_stones_with_one_bird",
				"can_hear_sign_language",
				"knows_victorias_secret",
			},
		},
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.User.Create(u)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", u)
	})

	t.Run("Read", func(t *testing.T) {
		u, err = m.User.Read(u.ID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", u)
	})

	t.Run("List", func(t *testing.T) {
		us, err := m.User.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", us)
	})

	t.Run("Update", func(t *testing.T) {
		uu := &User{
			Connection: "Username-Password-Authentication",
			Password:   "I don't need one",
		}
		err = m.User.Update(u.ID, uu)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", uu)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.User.Delete(u.ID)
		if err != nil {
			t.Error(err)
		}
	})
}
