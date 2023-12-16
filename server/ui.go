package server

import (
	"github.com/labstack/echo/v4"
	"lugosi/ui"
	"net/http"
)

func getContentHandler() http.Handler {
	content, _ := ui.Content()
	return http.FileServer(http.FS(content))
}

func UiRoutes(e *echo.Echo) {
	uiContentHandler := getContentHandler()
	e.GET("*", echo.WrapHandler(uiContentHandler))

	// page aliases
	e.GET("settings", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/settings.html")
	})
	e.GET("legacy", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/legacy.html")
	})
}
