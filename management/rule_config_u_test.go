package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitRuleConfig(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {

		ruleConfigInJson := `{
			"key": "key",
			"value": "value"
		}`

		decoder := json.NewDecoder(strings.NewReader(ruleConfigInJson))
		decoder.DisallowUnknownFields()
		var ruleConfig RuleConfig
		err = decoder.Decode(&ruleConfig)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		ruleConfig := RuleConfig{}
		bytes, _ := json.Marshal(ruleConfig)
		ruleConfigInJson := string(bytes)
		wanted := `{}`
		if ruleConfigInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", ruleConfigInJson, wanted)
		}
	})
}
