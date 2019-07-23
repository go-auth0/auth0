package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitRule(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {

		ruleInJson := `{
			"id": "generated id",
			"name": "name",
			"script": "script",
			"order": 123,
			"enabled": true
		}`

		decoder := json.NewDecoder(strings.NewReader(ruleInJson))
		decoder.DisallowUnknownFields()
		var rule Rule
		err = decoder.Decode(&rule)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		rule := Rule{}
		bytes, _ := json.Marshal(rule)
		ruleInJson := string(bytes)
		wanted := `{}`
		if ruleInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", ruleInJson, wanted)
		}
	})
}
