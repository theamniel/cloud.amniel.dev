package utils

import (
	"regexp"
	"strings"
)

var (
	supportedImage *regexp.Regexp = regexp.MustCompile(`^image\/(png|jpe?g|gif|webp)$`)
	supportedVideo *regexp.Regexp = regexp.MustCompile(`^video\/(mp4|m4v|webm|3gpp|quicktime)$`)
	isImage        *regexp.Regexp = regexp.MustCompile(`\.?(png|jpe?g|gif|webp)$`)
	isVideo        *regexp.Regexp = regexp.MustCompile(`\.?(mp4|m4v|webm|mov|3gpp?)$`)
)

// ParseFilename returns the file name and extension from the provided string.
func ParseFilename(raw string) (string, string) {
	values := strings.Split(raw, ".")
	n := len(values) - 1
	return strings.Join(values[:n], ""), values[n]
}

// SupportedMediaType returns true if the provided MIME type is supported.
func SupportedMediaType(mime string) bool {
	return supportedImage.MatchString(mime) ||
		supportedVideo.MatchString(mime) ||
		isVideo.MatchString(mime) ||
		isImage.MatchString(mime)
}

// IsVideo returns true if the provided string is a video file.
func IsVideo(raw string) bool {
	return isVideo.MatchString(raw)
}

// IsImage returns true if the provided string is an image file.
func IsImage(raw string) bool {
	return isImage.MatchString(raw)
}
