package storage

import (
	"encoding/json"
	"fmt"
	"happcmd/internal/profile"
	"os"
	"path/filepath"
	"strings"
)

type Storage struct {
	basePath string
}

func New(basePath string) *Storage {
	return &Storage{basePath: basePath}
}

func (s *Storage) profilesDir() string {
	return filepath.Join(s.basePath, ".happcmd", "profiles")
}

func (s *Storage) profilePath(name string) string {
	return filepath.Join(s.profilesDir(), name+".json")
}

func (s *Storage) Save(name string, p *profile.Profile) error {
	dir := s.profilesDir()

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	tmp := filepath.Join(dir, name+".json.tmp")
	final := s.profilePath(name)

	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, final)
}

func (s *Storage) Load(name string) (*profile.Profile, error) {
	data, err := os.ReadFile(s.profilePath(name))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("profile not found: %s", name)
		}
		return nil, err
	}

	p := &profile.Profile{}
	if err := json.Unmarshal(data, p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Storage) List() ([]string, error) {
	dir := s.profilesDir()

	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var res []string
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			res = append(res, strings.TrimSuffix(f.Name(), ".json"))
		}
	}

	return res, nil
}

func (s *Storage) Delete(name string) error {
	err := os.Remove(s.profilePath(name))
	if os.IsNotExist(err) {
		return fmt.Errorf("profile not found: %s", name)
	}
	return err
}

func (s *Storage) Exists(name string) (bool, error) {
	_, err := os.Stat(s.profilePath(name))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
