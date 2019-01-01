package services

import (
	"testing"

	"github.com/infernalslam/mock-api-server/services"
)

// type restApisImpMock interface {
// 	Heath() string
// }

func TestHeath(t *testing.T) {
	t.Run("saying hello to response", func(t *testing.T) {
		got := services.Heath()
		want := "Hello wolrd $$!@%#$@!#"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
