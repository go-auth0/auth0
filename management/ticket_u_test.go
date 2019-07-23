package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitTicket(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {

		ticketInJson := `{
			"result_url": "result url",
			"user_id": "user id",
			"ttl_sec": 123,
			"connection_id": "connection id",
			"email": "email",
			"ticket": "ticket"
		}`
		decoder := json.NewDecoder(strings.NewReader(ticketInJson))
		decoder.DisallowUnknownFields()
		var ticket Ticket
		err = decoder.Decode(&ticket)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		ticket := Ticket{}
		bytes, _ := json.Marshal(ticket)
		ticketInJson := string(bytes)
		wanted := `{}`
		if ticketInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", ticketInJson, wanted)
		}
	})
}
