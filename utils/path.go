package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPath(p string) (string, error) {
	if p == "" {
		return "", fmt.Errorf("empty path provided")
	}

	d, _, err := Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	d = filepath.Join(d, p)
	return d, nil
}

func GetOrCreatePath(p string) (string, error) {
	d, err := GetPath(p)
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(d); os.IsNotExist(err) {
		if err = os.MkdirAll(d, 0777); err != nil {
			return "", fmt.Errorf("failed to create directory: %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("failed to check directory existence: %w", err)
	}
	return d, nil
}
