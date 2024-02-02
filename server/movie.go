package server

import (
	"encoding/json"
	"io"
	"net/http"
	"videoteque/movie"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
	v := movie.VideoRef

	w.Header().Add("Cache-Control", "no-store")

	mime := getMime(v.Path())
	w.Header().Set("Content-Type", mime)

	reader := v.Reader()
	io.Copy(w, reader)
	defer reader.Close()
}

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	meta := movie.MetadataRef

	w.Header().Add("Cache-Control", "no-store")

	if meta == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(meta)
	}
}
