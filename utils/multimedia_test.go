package utils

import "testing"

func TestParseFilename(t *testing.T) {
	raw := "image_date_format.jpg"
	name, ext := ParseFilename(raw)
	if name != "image_date_format" || ext != "jpg" {
		t.Errorf("Expected image.jpg, got %s.%s", name, ext)
	}
}

var (
	supportedMediaTypes = []string{
		"image/png",
		"image/jpeg",
		"image/jpg",
		"image/webp",
		"image/gif",
		"video/mp4",
		"video/m4v",
		"video/webm",
		"video/3gpp",
		"video/quicktime",
	}
	supportedImages = []string{
		"png",
		"jpeg",
		"jpg",
		"webp",
		"gif",
	}
	supportedVideos = []string{
		"mp4",
		"m4v",
		"webm",
		"3gpp",
		"mov",
	}
)

func TestIsVideo(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		for _, mediaType := range supportedVideos {
			if !IsVideo(mediaType) {
				t.Errorf("Expected %s to be supported", mediaType)
			}
		}
	})

	t.Run("false", func(t *testing.T) {
		for _, mediaType := range supportedVideos {
			mediaType += "1"
			if IsVideo(mediaType) {
				t.Errorf("Expected false in %s", mediaType)
			}
		}
	})
}

func TestIsImage(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		for _, mediaType := range supportedImages {
			if !IsImage(mediaType) {
				t.Errorf("Expected %s to be supported", mediaType)
			}
		}
	})

	t.Run("false", func(t *testing.T) {
		for _, mediaType := range supportedImages {
			mediaType += "1"
			if IsImage(mediaType) {
				t.Errorf("Expected false in %s", mediaType)
			}
		}
	})
}

func TestSupportedMediaType(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		for _, mediaType := range supportedMediaTypes {
			if !SupportedMediaType(mediaType) {
				t.Errorf("Expected %s to be supported", mediaType)
			}
		}
	})

	t.Run("false", func(t *testing.T) {
		for _, mediaType := range supportedMediaTypes {
			mediaType += "1"
			if SupportedMediaType(mediaType) {
				t.Errorf("Expected false in %s", mediaType)
			}
		}
	})
}
