package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitResourceServer(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {

		resourceServerInJson := `{
			"id": "generated id",
			"name": "name",
			"identifier": "identifier",
			"scopes": [
				{
					"value": "scope",
					"description": "scope description"
				}
			],
			"signing_alg": "RS256",
			"signing_secret": "signing secret",
			"allow_offline_access": true,
			"token_lifetime": 123,
			"token_lifetime_for_web": 123,
			"skip_consent_for_verifiable_first_party_clients": true,
			"verificationLocation": "URI to retrieve JWKs",
			"options": {
				"key": "value"
			},
			"enforce_policies": true
		}`

		decoder := json.NewDecoder(strings.NewReader(resourceServerInJson))
		decoder.DisallowUnknownFields()
		var resourceServer ResourceServer
		err = decoder.Decode(&resourceServer)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		resourceServer := ResourceServer{}
		bytes, _ := json.Marshal(resourceServer)
		resourceServerInJson := string(bytes)
		wanted := `{}`
		if resourceServerInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", resourceServerInJson, wanted)
		}
	})
}
