package expect

import (
	"reflect"
	"testing"
)

// Expect performs an assertion that expects x to equal y. The assertion is
// performed using reflect.DeepEqual.
func Expect(t *testing.T, x, y interface{}) bool {
	t.Helper()
	if !reflect.DeepEqual(x, y) {
		t.Errorf("Expected %v to equal %v\n", x, y)
		return false
	}
	return true
}
