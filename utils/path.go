package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetPath returns the absolute path by joining the executable directory with the provided path.
func GetPath(p ...string) (string, error) {
	if len(p) == 0 {
		return "", fmt.Errorf("empty path provided")
	}

	d := filepath.Join(p...)
	return d, nil
}

// GetOrCreatePath returns the absolute path and creates the directory if it does not exist.
func GetOrCreatePath(p string) (string, error) {
	d, err := GetPath(p)
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(d); os.IsNotExist(err) {
		if err = os.MkdirAll(d, 0755); err != nil {
			return "", fmt.Errorf("failed to create directory: %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("failed to check directory existence: %w", err)
	}
	return d, nil
}
