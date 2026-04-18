package profile

import "testing"

func TestNewProfile_Ok(t *testing.T) {
	// Arrange
	t.Parallel()
	p := NewProfile("Test", nil, nil, nil)

	// Assert
	if p == nil {
		t.Fatal("NewProfile returned nil")
	}

	if p.Name != "Test" {
		t.Error("NewProfile returned wrong name")
	}

	if len(p.DirectSites) == 0 {
		t.Error("expected DirectSites to be non-empty")
	}

	if len(p.BlockSites) == 0 {
		t.Error("expected BlockSites to be non-empty")
	}
}
