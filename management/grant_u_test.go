package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitGrant(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		grantInJson := `{
			"id": "generated id",
			"clientID": "id of the client",
			"user_id": "id of the user",
			"audience": "audience",
			"scope": [
				"scope_1",
				"scope_2"
			]
		}`

		decoder := json.NewDecoder(strings.NewReader(grantInJson))
		decoder.DisallowUnknownFields()
		var grant Grant
		err = decoder.Decode(&grant)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		grant := Grant{}
		bytes, _ := json.Marshal(grant)
		grantInJson := string(bytes)
		wanted := `{"user_id":null}`
		if grantInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", grantInJson, wanted)
		}
	})
}
