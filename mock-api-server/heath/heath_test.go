package heath

import (
	"testing"

	"github.com/infernalslam/mock-api-server/heath"
)

func TestHeath(t *testing.T) {
	t.Run("saying hello to response", func(t *testing.T) {
		got := heath.SayHello()
		want := "Hello wolrd $$!@%#$@!#"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
