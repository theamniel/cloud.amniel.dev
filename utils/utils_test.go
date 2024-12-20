package utils_test

import (
	"testing"

	"cloud.amniel.dev/utils"
)

func TestExecutable(t *testing.T) {
	t.Run("TestExecutable", func(t *testing.T) {
		t.Parallel()
		_, _, err := utils.Executable()
		if err != nil {
			t.Errorf("TestExecutable failed: %v", err)
		}
	})
}

func TestRandomString(t *testing.T) {
	t.Run("TestRandomString", func(t *testing.T) {
		t.Parallel()
		s := utils.RandomString(10)
		if len(s) != 10 {
			t.Errorf("TestRandomString failed: %v", s)
		}
	})
}
