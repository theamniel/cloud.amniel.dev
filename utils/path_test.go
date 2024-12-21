package utils

import (
	"os"
	"path/filepath"
	"testing"
)

// TestGetPath tests the GetPath function.
func TestGetPath(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	path, err := GetPath(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the path exists
	if _, err = os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("path does not exist: %s", path)
	}

	// Check if the path is a directory
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatalf("path is not a directory: %s", path)
	}
}

// TestGetOrCreatePath tests the GetOrCreatePath function.
func TestGetOrCreatePath(t *testing.T) {
	// Create a temporary directory
	dir, err := GetOrCreatePath("temp")
	if err != nil {
		t.Fatal(err)
	}

	// Check if the directory exists
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		t.Fatalf("directory does not exist: %s", dir)
	}

	// Check if the directory is indeed a directory
	info, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatalf("path is not a directory: %s", dir)
	}

	// Check if the directory is writable
	testFile := filepath.Join(dir, "testfile")
	if err = os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("directory is not writable: %s", dir)
	}
	if err = os.Remove(testFile); err != nil {
		t.Fatalf("failed to remove test file: %s", testFile)
	}

	// Remove the temporary directory
	if err = os.Remove(dir); err != nil {
		t.Fatalf("failed to remove directory: %s", dir)
	}
}

// TestGetOrCreatePathExisting tests the GetOrCreatePath function with an existing directory.
func TestGetOrCreatePathExisting(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Use the existing directory
	dir, err := GetOrCreatePath(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the directory exists
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		t.Fatalf("directory does not exist: %s", dir)
	}

	// Check if the directory is indeed a directory
	info, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatalf("path is not a directory: %s", dir)
	}
}
