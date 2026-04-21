package storage

import (
	"encoding/json"
	"happcmd/internal/profile"
	"os"
	"path/filepath"
	"strings"
)

func Save(name string, profile *profile.Profile) error {
	path, err := profilesDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	p, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(path, name+".json"), p, 0644)
}

func Load(name string) (*profile.Profile, error) {
	path, err := profilesDir()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath.Join(path, name+".json"))
	if err != nil {
		return nil, err
	}
	p := &profile.Profile{}
	err = json.Unmarshal(data, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func List() ([]string, error) {
	path, err := profilesDir()
	if err != nil {
		return nil, err
	}
	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var res []string
	for _, f := range files {
		res = append(res, strings.TrimSuffix(f.Name(), ".json"))
	}
	return res, nil
}

func Delete(name string) error {
	path, err := profilesDir()
	if err != nil {
		return err
	}
	return os.Remove(filepath.Join(path, name+".json"))
}

func Exists(name string) (bool, error) {
	path, err := profilesDir()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(filepath.Join(path, name+".json"))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func profilesDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".happcmd", "profiles"), nil
}
