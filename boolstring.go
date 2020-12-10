package auth0

import (
	"bytes"
	"encoding/json"
	"errors"
)

var (
	errUnmarshalBoolString = errors.New("BoolString: json.Unmarshal failed")

	trueStringJSON  = json.RawMessage(`"true"`)
	falseStringJSON = json.RawMessage(`"false"`)

	trueJSON  = json.RawMessage(`true`)
	falseJSON = json.RawMessage(`false`)
	nullJSON  = json.RawMessage(`null`)
)

type BoolString bool

func NewBoolString(v bool) *BoolString {
	result := BoolString(v)
	return &result
}

func (b BoolString) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

func (b *BoolString) UnmarshalJSON(data []byte) error {
	var v json.RawMessage
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch {
	case bytes.Equal(v, trueStringJSON), bytes.Equal(v, trueJSON):
		*b = true
	case bytes.Equal(v, falseStringJSON), bytes.Equal(v, falseJSON), bytes.Equal(v, nullJSON):
		*b = false
	default:
		return errUnmarshalBoolString
	}

	return nil
}
