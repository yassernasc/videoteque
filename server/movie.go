package server

import (
	"encoding/json"
	"io"
	"net/http"
	"videoteque/movie"
	"videoteque/torrent"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
	format := movie.Video.Format
	content := movie.Video.Payload

	switch format {
	case movie.Magnet:
		stream, displayPath := torrent.Stream(content)
		w.Header().Set("Content-Type", getMime(displayPath))
		io.Copy(w, stream)
	case movie.File:
		http.ServeFile(w, r, content)
	case movie.Url:
		http.Redirect(w, r, content, http.StatusMovedPermanently)
	}
}

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	m := movie.Video.Metadata

	if m == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	}
}
