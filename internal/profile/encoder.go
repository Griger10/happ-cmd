package profile

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Mode string

const (
	ModeAdd   Mode = "add"
	ModeOnAdd Mode = "onadd"
)

func (p *Profile) ToURL(mode Mode) (string, error) {
	j, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	s := base64.StdEncoding.EncodeToString(j)
	return fmt.Sprintf("happ://routing/%s/%s", mode, s), nil
}
