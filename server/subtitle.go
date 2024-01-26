package server

import (
	"fmt"
	"net/http"
	"videoteque/subtitle"
)

func subtitleHandler(w http.ResponseWriter, r *http.Request) {
	s := subtitle.Get()

	w.Header().Add("Cache-Control", "no-store")

	if s == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "text/vtt")
		fmt.Fprint(w, s)
	}
}
