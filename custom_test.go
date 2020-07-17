package auth0

import "testing"

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

