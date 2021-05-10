package management

import (
	"testing"
)

func TestAnomaly(t *testing.T) {

	t.Run("CheckIP", func(t *testing.T) {
		err := m.Anomaly.CheckIP("1.1.1.1")
		if err != nil {
			t.Error(err)
		}
	})

}
