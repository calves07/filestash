package plg_video_thumbnail

import "testing"

func TestSupportedVideoThumbnailMIMETypes(t *testing.T) {
	expected := []string{
		"video/mp4",
		"video/x-matroska",
		"video/x-msvideo",
		"video/quicktime",
		"video/webm",
		"video/x-m4v",
		"video/m4v",
	}

	if len(supportedVideoThumbnailMIMETypes) != len(expected) {
		t.Fatalf("unexpected mime count: got=%d want=%d", len(supportedVideoThumbnailMIMETypes), len(expected))
	}

	seen := map[string]bool{}
	for _, mime := range supportedVideoThumbnailMIMETypes {
		seen[mime] = true
	}

	for _, mime := range expected {
		if !seen[mime] {
			t.Fatalf("missing mime type: %s", mime)
		}
	}
}
