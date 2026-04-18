package profile

import "testing"

func TestDecodeUrl(t *testing.T) {
	// Arrange
	t.Parallel()
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid happ url",
			url:     mustEncodeProfile(t),
			wantErr: false,
		},
		{
			name:    "invalid base64",
			url:     "happ://routing/add/!!!invalid!!!",
			wantErr: true,
		},
		{
			name:    "empty string",
			url:     "",
			wantErr: true,
		},
		{
			name:    "invalid url format",
			url:     "not-a-happ-url",
			wantErr: true,
		},
	}

	// Act + Assert
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := DecodeUrl(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func mustEncodeProfile(t *testing.T) string {
	// Arrange
	t.Helper()
	p := NewProfile("Test", nil, nil, nil)
	// Act
	url, err := p.ToURL(ModeAdd)
	// Assert
	if err != nil {
		t.Fatalf("failed to encode profile: %v", err)
	}
	return url
}
