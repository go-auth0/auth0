package auth0

import (
	"encoding/json"
	"testing"
)


func TestConvertibleBool(t *testing.T) {
	for _, test := range []struct {
		in       *ConvertibleBoolean
		expected ConvertibleBoolean
	}{
		{nil, false},
		{ConvertibleBool(false), false},
		{ConvertibleBool(true), true},
	} {
		have := ConvertibleBoolValue(test.in)
		if have != test.expected {
			t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
		}
	}
}

func TestConvertibleBoolean_UnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		in string
		expected *ConvertibleBoolean
	}{
		{`{"bool": false}`, ConvertibleBool(false)},
		{`{"bool": "false"}`, ConvertibleBool(false)},
		{`{"bool": true}`, ConvertibleBool(true)},
		{`{"bool": "true"}`, ConvertibleBool(true)},
	} {
		var ts struct {
			Bool *ConvertibleBoolean `json:"bool,omitempty"`
		}
		err := json.Unmarshal([]byte(test.in),&ts)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if ConvertibleBoolValue(ts.Bool) != ConvertibleBoolValue(test.expected) {
			t.Errorf("unexpected output. have %v, expected %v", ts.Bool, test.expected)
		}
	}


}

