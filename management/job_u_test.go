package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitJob(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		jobInJson := `{
			"id": "generated id",
			"status": "pending",
			"type": "type",
			"created_at": "2006-01-02T15:04:05+02:00",
			"user_id": "id of the user",
			"client_id": "id of the client",
			"connection_id": "id of the connection",
			"location": "url to download the result of the job",
			"percentage_done": 58,
			"time_left_seconds": 123,
			"format": "csv",
			"time_left_seconds": 1000,
			"fields": [
				{
					"name": "identities[0].connection",
					"export_as": "provider"
				},
				{
					"name": "email"
				}
			],
			"users": [
				{
					"email": "john.doe@contoso.com",
					"email_verified": false,
					"app_metadata": {
						"roles": ["admin"],
						"plan": "premium"
					},
					"user_metadata": {
						"theme": "light"
					}
				}
			],
			"upsert": true,
			"external_id": "correlation id between jobs",
			"send_completion_email": false
		}`

		decoder := json.NewDecoder(strings.NewReader(jobInJson))
		decoder.DisallowUnknownFields()
		var job Job
		err = decoder.Decode(&job)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		job := Job{}
		bytes, _ := json.Marshal(job)
		jobInJson := string(bytes)
		wanted := `{}`
		if jobInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", jobInJson, wanted)
		}
	})
}
