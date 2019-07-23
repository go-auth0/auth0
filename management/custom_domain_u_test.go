package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitCustomDomain(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		customDomainInJson := `{
			"custom_domain_id": "custom domain id",
			"domain": "mycompany.org",
			"type": "auth0_managed_certs",
			"primary": true,
			"status": "ready",
			"verification_method": "txt",
			"verification": {
				"methods": [
					{
						"key": "value"
					}
				]
			}
		}`

		decoder := json.NewDecoder(strings.NewReader(customDomainInJson))
		decoder.DisallowUnknownFields()
		var customDomain CustomDomain
		err = decoder.Decode(&customDomain)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		customDomain := CustomDomain{}
		bytes, _ := json.Marshal(customDomain)
		customDmainInJson := string(bytes)
		wanted := "{}"
		if customDmainInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", customDmainInJson, wanted)
		}
	})
}
