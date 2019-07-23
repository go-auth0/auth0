package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitStat(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {

		dailyStatInJson := `{
			"date": "2006-01-02T15:04:05+02:00",
			"logins": 123,
			"signups": 123,
			"leaked_passwords": 123,
			"updated_at": "2006-01-02T15:04:05+02:00",
			"created_at": "2006-01-02T15:04:05+02:00"
		}`
		decoder := json.NewDecoder(strings.NewReader(dailyStatInJson))
		decoder.DisallowUnknownFields()
		var dailyStat DailyStat
		err = decoder.Decode(&dailyStat)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		dailyStat := DailyStat{}
		bytes, _ := json.Marshal(dailyStat)
		dailyStatInJson := string(bytes)
		wanted := `{"date":null,"logins":null,"signups":null,"leaked_passwords":null,"updated_at":null,"created_at":null}`
		if dailyStatInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", dailyStatInJson, wanted)
		}
	})
}
