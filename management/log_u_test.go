package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitLog(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {

		logInJson := `{
			"_id": "generated id",
			"log_id": "log id",
			"date": "2006-01-02T15:04:05+02:00",
			"type": "type",
			"client_id": "id of the client",
			"client_name": "id of the client",
			"ip": "IP of the log event source",
			"location_info": {
				"key": "value"
			},
			"details": {
				"key": "value"
			},
			"user_id": "id of the user"
		}`
		decoder := json.NewDecoder(strings.NewReader(logInJson))
		decoder.DisallowUnknownFields()
		var job Log
		err = decoder.Decode(&job)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		log := Log{}
		bytes, _ := json.Marshal(log)
		logInJson := string(bytes)
		wanted := `{"_id":null,"log_id":null,"date":null,"type":null,"client_id":null,"client_name":null,"ip":null,"location_info":null,"details":null,"user_id":null}`
		if logInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", logInJson, wanted)
		}
	})
}
