package uuid

import (
	"testing"
)

func TestDefaultGenerator(t *testing.T) {
	t.Run("DefaultGeneratorが正常に動作する", func(t *testing.T) {
		// Test that we can generate a UUID
		uuid, err := DefaultGenerator()
		if err != nil {
			t.Errorf("DefaultGenerator() error = %v, want nil", err)
			return
		}
		if uuid.String() == "" {
			t.Errorf("DefaultGenerator() generated empty UUID")
		}
	})
}