package utils

import (
	"os"
	"testing"
)

// TestGetPath tests the GetPath function.
func TestGetPath(t *testing.T) {
	// Get the executable path
	tempDir := t.TempDir()
	path, err := GetPath(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the path exists
	if _, err = os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("path does not exist: %s", path)
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

	// Remove the temporary directory
	if err = os.Remove(dir); err != nil {
		t.Fatalf("failed to remove directory: %s", dir)
	}
}
