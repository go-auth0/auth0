package custom

import (
	"bytes"
	"encoding/json"
	"gopkg.in/auth0.v4"
	"testing"
)

func TestConvertibleBoolean_UnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		in       string
		expected *ConvertibleBoolean
	}{
		{`{"bool": false}`, auth0.ConvertibleBool(false)},
		{`{"bool": "false"}`, auth0.ConvertibleBool(false)},
		{`{"bool": true}`, auth0.ConvertibleBool(true)},
		{`{"bool": "true"}`, auth0.ConvertibleBool(true)},
	} {
		var ts struct {
			Bool *ConvertibleBoolean `json:"bool,omitempty"`
		}
		err := json.Unmarshal([]byte(test.in), &ts)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if auth0.ConvertibleBoolValue(ts.Bool) != auth0.ConvertibleBoolValue(test.expected) {
			t.Errorf("unexpected output. have %v, expected %v", ts.Bool, test.expected)
		}
	}
}

func TestConvertibleBoolean_MarshalJSON(t *testing.T) {
	for _, test := range []struct {
		in       *ConvertibleBoolean
		expected []byte
	}{
		{auth0.ConvertibleBool(false), []byte(`{"bool":false}`)},
		{auth0.ConvertibleBool(true), []byte(`{"bool":true}`)},
	} {
		var ts struct {
			Bool *ConvertibleBoolean `json:"bool,omitempty"`
		}
		ts.Bool = test.in
		str, err := json.Marshal(&ts)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if !bytes.Equal(str, test.expected) {
			t.Errorf("unexpected output. have %v, expected %v", ts.Bool, test.expected)
		}
	}
}
