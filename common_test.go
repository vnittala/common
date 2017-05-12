package common

import (
	"testing"
)

func TestNewUUID(t *testing.T) {

	id := NewUUID()

	if len(id) < 1 {

		t.Errorf("Generated ID is null, got: %d", id)
	}

}
