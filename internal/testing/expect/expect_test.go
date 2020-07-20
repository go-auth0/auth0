package expect

import "testing"

func TestExpect(t *testing.T) {
	for _, test := range []struct {
		a, b interface{}
	}{
		{"", ""},
		{"a", "a"},
		{1, 1},
		{[]string{"1"}, []string{"1"}},
		{[]int{1}, []int{1}},
	} {
		Expect(t, test.a, test.b)
	}
}
