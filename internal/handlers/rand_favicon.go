package handlers

import (
	"image"
	"net/http"
)

func FaviconGen(w http.ResponseWriter, r *http.Request) {
	image.NewRGBA(image.Rect(0, 0, 48, 48))
}
