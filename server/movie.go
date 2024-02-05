package server

import (
	"encoding/json"
	"net/http"
	"time"
	"videoteque/movie"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
	v := movie.VideoRef

	w.Header().Add("Cache-Control", "no-store")
	http.ServeContent(w, r, v.Path(), time.Now(), v.Reader())
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
