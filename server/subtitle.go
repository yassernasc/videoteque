package server

import (
	"fmt"
	"net/http"
	"videoteque/subtitle"
)

func subtitleHandler(w http.ResponseWriter, r *http.Request) {
	sub, code, err := subtitle.Get()

	w.Header().Add("Cache-Control", "no-store")

	if code == http.StatusOK {
		w.Header().Set("Content-Type", "text/vtt")
		fmt.Fprint(w, sub)
	} else {
		if err != nil {
			fmt.Fprint(w, err)
		}
		w.WriteHeader(code)
	}
}
