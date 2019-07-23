package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitUser(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {

		userInJson := `{
			"user_id": "user_id",
			"connection": "connection",
			"email": "email",
			"name": "name",
			"given_name": "given name",
			"family_name": "family name",
			"username": "username",
			"nickname": "nickname",
			"password": "password",
			"phone_number": "phone_number",
			"created_at": "2006-01-02T15:04:05+02:00",
			"updated_at": "2006-01-02T15:04:05+02:00",
			"last_login": "2006-01-02T15:04:05+02:00",
			"user_metadata": {
				"key": "value"
			},
			"identities": [
				{
					"connection": "connection",
					"user_id": "user id",
					"provider": "auth0",
					"isSocial": false
				}
			],
			"email_verified": true,
			"verify_email": true,
			"phone_verified": true,
			"app_metadata": {
				"key": "value"
			},
			"picture": "picture"
		}`

		decoder := json.NewDecoder(strings.NewReader(userInJson))
		decoder.DisallowUnknownFields()
		var user User
		err = decoder.Decode(&user)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields for Role", func(t *testing.T) {
		user := User{}
		bytes, _ := json.Marshal(user)
		userInJson := string(bytes)
		wanted := `{}`
		if userInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", userInJson, wanted)
		}
	})
}
