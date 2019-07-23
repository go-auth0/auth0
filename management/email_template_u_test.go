package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitEmailTemplate(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		emailTemplateInJson := `{
			"template": "ses",
			"body": "html page",
			"from": "support@mycompany.com",
			"resultUrl": "https://mycompany.org",
			"subject": "subject",
			"syntax": "syntax",
			"urlLifetimeInSeconds": 123,
			"enabled": true
		}`

		decoder := json.NewDecoder(strings.NewReader(emailTemplateInJson))
		decoder.DisallowUnknownFields()
		var emailTemplate EmailTemplate
		err = decoder.Decode(&emailTemplate)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		emailTemplate := EmailTemplate{}
		bytes, _ := json.Marshal(emailTemplate)
		emailTemplateInJson := string(bytes)
		wanted := `{"template":null,"enabled":null}`
		if emailTemplateInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", emailTemplateInJson, wanted)
		}
	})
}
