package profile

type Preset struct {
	Name        string
	Description string
	Factory     func(name string) *Profile
}

var Presets = map[string]Preset{
	"ru_default": {
		Name:        "Russia Default",
		Description: "Russian sites direct, ads blocked, everything else proxied",
		Factory:     func(name string) *Profile { return NewProfile(name, nil, nil, nil) },
	},
	"ru_strict": {
		Name:        "Russia Strict",
		Description: "Only local networks direct, everything else proxied",
		Factory:     NewStrictProfile,
	},
	"bypass_all": {
		Name:        "Bypass All",
		Description: "Everything through proxy, no exceptions",
		Factory:     NewBypassAllProfile,
	},
}
