package helpers

import (
	"net/http"
)

func InfoExtention(data []byte) string {

	mimeType := http.DetectContentType(data)

	switch mimeType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"
	case "video/mp4":
		return ".mp4"
	default:
		return ""
	}
}
