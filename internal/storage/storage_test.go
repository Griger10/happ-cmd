package storage

import (
	"happcmd/internal/profile"
	"testing"
)

func TestStorage_SaveLoad(t *testing.T) {
	// Arrange
	t.Parallel()

	dir := t.TempDir()
	s := New(dir)

	name := "test"
	p := profile.NewProfile(name, nil, nil, nil)

	// Act + Assert
	err := s.Save(name, p)
	if err != nil {
		t.Fatalf("save failed: %v", err)
	}

	loaded, err := s.Load(name)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if loaded.Name != p.Name {
		t.Fatalf("expected name %s, got %s", p.Name, loaded.Name)
	}
}

func TestStorage_Load_NotFound(t *testing.T) {
	t.Parallel()

	s := New(t.TempDir())

	_, err := s.Load("unknown")
	if err == nil {
		t.Fatal("expected error for non-existing profile")
	}
}

func TestStorage_Exists(t *testing.T) {
	// Arrange
	t.Parallel()

	s := New(t.TempDir())
	name := "test"

	// Act + Assert
	exists, err := s.Exists(name)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("should not exist")
	}

	if err := s.Save(name, profile.NewProfile(name, nil, nil, nil)); err != nil {
		t.Fatal(err)
	}

	exists, err = s.Exists(name)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("should exist")
	}
}

func TestStorage_List(t *testing.T) {
	// Arrange
	t.Parallel()

	s := New(t.TempDir())
	if err := s.Save("test1", profile.NewProfile("test1", nil, nil, nil)); err != nil {
		t.Fatal(err)
	}
	if err := s.Save("test2", profile.NewProfile("test2", nil, nil, nil)); err != nil {
		t.Fatal(err)
	}

	// Act + Assert
	list, err := s.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 2 {
		t.Fatal("expected 2 profiles, got", len(list))
	}
}

func TestStorage_Delete(t *testing.T) {
	// Arrange
	t.Parallel()

	s := New(t.TempDir())
	name := "test"

	if err := s.Save(name, profile.NewProfile(name, nil, nil, nil)); err != nil {
		t.Fatal(err)
	}

	// Act + Assert
	err := s.Delete(name)
	if err != nil {
		t.Fatal(err)
	}

	exists, err := s.Exists(name)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("profile should be deleted")
	}
}
