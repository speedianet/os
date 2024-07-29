package valueObject

import (
	"testing"
)

func TestVirtualHostType(t *testing.T) {
	t.Run("ValidVirtualHostType", func(t *testing.T) {
		validVhostTypes := []interface{}{
			"primary",
			"top-level",
			"subdomain",
			"wildcard",
			"alias",
		}

		for _, vhostType := range validVhostTypes {
			_, err := NewVirtualHostType(vhostType)
			if err != nil {
				t.Errorf("Expected no error for %s, got %v", vhostType, err)
			}
		}
	})

	t.Run("InvalidVirtualHostType", func(t *testing.T) {
		invalidVhostTypes := []string{
			"extradomain",
			"low-level",
			"secondary",
			"legacy",
		}

		for _, vhostType := range invalidVhostTypes {
			_, err := NewVirtualHostType(vhostType)
			if err == nil {
				t.Errorf("Expected error for %s, got nil", vhostType)
			}
		}
	})
}