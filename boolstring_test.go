package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestBoolStringMarshal(t *testing.T) {
	tests := []struct {
		input BoolString
		want  []byte
	}{
		{
			input: BoolString(true),
			want:  []byte(`true`),
		},
		{
			input: BoolString(false),
			want:  []byte(`false`),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input=%v", test.input), func(t *testing.T) {
			got, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(test.want, got) {
				t.Fatalf("wanted %q, got %q", test.want, got)
			}
		})
	}

}
func TestBoolStringUnmarshal(t *testing.T) {
	var doesNotMatter json.RawMessage
	errUnexpectedEndOfInput := json.Unmarshal([]byte(""), &doesNotMatter)

	tests := []struct {
		input     []byte
		want      BoolString
		wantError error
	}{
		{
			input: trueJSON,
			want:  BoolString(true),
		},
		{
			input: trueStringJSON,
			want:  BoolString(true),
		},
		{
			input: falseJSON,
			want:  BoolString(false),
		},
		{
			input: falseStringJSON,
			want:  BoolString(false),
		},
		{
			input: nullJSON,
			want:  BoolString(false),
		},
		{
			input:     []byte(`1`),
			wantError: errUnmarshalBoolString,
		},
		{
			input:     []byte(`0`),
			wantError: errUnmarshalBoolString,
		},
		{
			input:     []byte(``),
			wantError: errUnexpectedEndOfInput,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input=%s", test.input), func(t *testing.T) {
			var got BoolString

			if err := json.Unmarshal(test.input, &got); !reflect.DeepEqual(test.wantError, err) {
				t.Fatalf("wanted err: %v, got %v", test.wantError, err)
			}

			if test.want != got {
				t.Fatalf("wanted %v, got %v", test.want, got)
			}
		})
	}
}
