package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

const characters string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var (
	bigint = big.NewInt(int64(len(characters)))
)

// Executable returns the directory and file name of the executable, excluding the ".exe" extension if present.
func Executable() (string, string, error) {
	e, err := os.Executable()
	if err != nil {
		return "", "", fmt.Errorf("failed to get executable path: %w", err)
	}
	if e == "" {
		return "", "", fmt.Errorf("empty executable path")
	}
	d, f := filepath.Split(e)

	if f == "" {
		return "", "", fmt.Errorf("empty file name")
	}
	if ext := filepath.Ext(f); ext == ".exe" {
		f = f[:len(f)-len(ext)]
	}
	return d, f, nil
}

// RandomString generates a random string of the specified length.
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, bigint)
		if err != nil {
			panic(err)
		}
		b[i] = characters[num.Int64()]
	}
	return string(b)
}

// Hash returns the SHA-1 hash of the input byte slice.
func Hash(b []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(b))
}

// Base64 returns the base64 encoding of the input byte slice.
func Base64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
