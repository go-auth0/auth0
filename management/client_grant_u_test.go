package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitClientGrant(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		clientGrantInJson := `{
			"id": "generated_id",
			"client_id": "client_id",
			"audience": "audience",
			"scope": [
				"scope_1",
				"scope_2"
			]
		}`

		decoder := json.NewDecoder(strings.NewReader(clientGrantInJson))
		decoder.DisallowUnknownFields()
		var clientGrant ClientGrant
		err = decoder.Decode(&clientGrant)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		clientGrant := ClientGrant{}
		bytes, _ := json.Marshal(clientGrant)
		clientGrantInJson := string(bytes)
		wanted := `{"scope":null}`
		if clientGrantInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", clientGrantInJson, wanted)
		}
	})
}
