package server

import (
	"net/http"
	"videoteque/ui"
)

func uiHandler() http.Handler {
	content, _ := ui.Content()
	return http.FileServer(http.FS(content))
}
