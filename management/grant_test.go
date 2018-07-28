package management

import (
	"testing"
)

func TestGrant(t *testing.T) {

	var err error

	t.Run("List", func(t *testing.T) {
		var gs []*Grant
		gs, err = m.Grant.List()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", gs)
	})
}
