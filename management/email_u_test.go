package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitEmail(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		emailInJson := `{
			"name": "ses",
			"enabled": true,
			"default_from_address": "support@mycompany.org",
			"credentials": {
				"api_user": "api user",
				"api_key": "api key",
				"accessKeyId": "access Key Id",
				"secretAccessKey": "secret Access Key",
				"region": "region",
				"smtp_host": "smtp host",
				"smtp_port": 1234,
				"smtp_user": "smtp user",
				"smtp_pass": "smtp pass"
			}, 
			"settings": {
				"key": "value"
			}
		}`

		decoder := json.NewDecoder(strings.NewReader(emailInJson))
		decoder.DisallowUnknownFields()
		var email Email
		err = decoder.Decode(&email)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		email := Email{}
		bytes, _ := json.Marshal(email)
		emailInJson := string(bytes)
		wanted := "{}"
		if emailInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", emailInJson, wanted)
		}
	})
}
