package profile

import "testing"

func TestPresets(t *testing.T) {
	tests := []struct {
		name       string
		factory    func(string) *Profile
		wantDirect bool
		wantBlock  bool
	}{
		{
			name:       "ru_default",
			factory:    func(n string) *Profile { return NewProfile(n, nil, nil, nil) },
			wantDirect: true,
			wantBlock:  true,
		},
		{
			name:       "ru_strict",
			factory:    NewStrictProfile,
			wantDirect: false,
			wantBlock:  true,
		},
		{
			name:       "bypass_all",
			factory:    NewBypassAllProfile,
			wantDirect: false,
			wantBlock:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := tt.factory("Test")
			if tt.wantDirect && len(p.DirectSites) == 0 {
				t.Error("expected DirectSites to be non-empty")
			}
			if !tt.wantDirect && len(p.DirectSites) != 0 {
				t.Error("expected DirectSites to be empty")
			}
			if tt.wantBlock && len(p.BlockSites) == 0 {
				t.Error("expected BlockSites to be non-empty")
			}
			if !tt.wantBlock && len(p.BlockSites) != 0 {
				t.Error("expected BlockSites to be empty")
			}
		})
	}
}
