package profile

import "testing"

func TestNewProfile_ToURL(t *testing.T) {
	// Arrange
	t.Parallel()
	p := NewProfile("Test", nil, nil, nil)

	// Act + Assert
	_, err := p.ToURL(ModeAdd)
	if err != nil {
		t.Fatal("ToURL returned unexpected error")
	}
}
