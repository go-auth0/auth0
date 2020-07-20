package custom

import "strconv"

// ConvertibleBoolean is a custom bool type that is able to unmarshal/marshal from both string and bool types
type ConvertibleBoolean bool

// UnmarshalJSON handles unmarshalling of the ConvertibleBoolean value in the case of string and also bool values
func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := trimQuotes(string(data))
	b, err := strconv.ParseBool(asString)
	if err != nil {
		return err
	}
	*bit = ConvertibleBoolean(b)
	return nil
}

// MarshalJSON handles marshalling of the ConvertibleBoolean value. We will always return the value without quotes.
func (bit ConvertibleBoolean) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatBool(bool(bit))), nil
}

func trimQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}
