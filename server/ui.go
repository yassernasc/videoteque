package server

import (
	"net/http"
	"videoteque/fs"
	"videoteque/ui"
)

var uiServer http.Handler

func init() {
	content, _ := ui.Content()
	uiServer = http.FileServer(http.FS(content))
}

func uiHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if isPage(path) {
		r.URL.Path = path + ".html"
	}

	if isStaticAsset(path) {
		w.Header().Add("Cache-Control", "max-age=31536000, immutable")
	}

	uiServer.ServeHTTP(w, r)
}

func isStaticAsset(path string) bool {
	exts := [...]string{".js", ".css", ".ttf"}
	pathExt := fs.Ext(path)

	for _, e := range exts {
		if pathExt == e {
			return true
		}
	}

	return false
}

func isPage(path string) bool {
	pages := [...]string{"/settings", "/legacy"}

	for _, page := range pages {
		if path == page {
			return true
		}
	}

	return false
}
