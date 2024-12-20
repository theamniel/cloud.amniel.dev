package utils_test

import (
	"testing"

	"cloud.amniel.dev/utils"
)

// TestHash tests the Hash function.
func TestHash(t *testing.T) {
	t.Run("TestHash", func(t *testing.T) {
		t.Parallel()
		s := utils.Hash([]byte("test"))
		if s == "" {
			t.Errorf("TestHash failed: %v", s)
		}
	})
}

// TestBase64 tests the Base64 function.
func TestBase64(t *testing.T) {
	t.Run("TestBase64", func(t *testing.T) {
		t.Parallel()
		s := utils.Base64([]byte("test"))
		if s == "" {
			t.Errorf("TestBase64 failed: %v", s)
		}
	})
}

// TestExecutable tests the Executable function.
func TestExecutable(t *testing.T) {
	t.Run("TestExecutable", func(t *testing.T) {
		t.Parallel()
		_, _, err := utils.Executable()
		if err != nil {
			t.Errorf("TestExecutable failed: %v", err)
		}
	})
}

// TestRandomString tests the RandomString function.
func TestRandomString(t *testing.T) {
	t.Run("TestRandomString", func(t *testing.T) {
		t.Parallel()
		s := utils.RandomString(10)
		if len(s) != 10 {
			t.Errorf("TestRandomString failed: %v", s)
		}
	})
}
