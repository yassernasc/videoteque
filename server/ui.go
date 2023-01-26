package server

import (
	"github.com/labstack/echo/v4"
	"log"
	"lugosi/ui"
	"net/http"
)

func getContentHandler() http.Handler {
	content, err := ui.Content()
	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(http.FS(content))
}

func UiRoutes(e *echo.Echo) {
	uiContentHandler := getContentHandler()
	e.GET("*", echo.WrapHandler(uiContentHandler))
}
