package server

import (
	"github.com/labstack/echo/v4"
	"lugosi/judgment"
	"lugosi/storage"
	"lugosi/torrent"
	"net/http"
)

func MovieRoutes(e *echo.Echo) {
	e.GET("/movie", func(c echo.Context) error {
		entry := storage.Movie()

		if judgment.IsUrl(entry) {
			return c.Redirect(http.StatusMovedPermanently, entry)
		} else if judgment.IsFile(entry) {
			return c.Redirect(http.StatusMovedPermanently, "/movie/static")
		} else {
			stream, mime := torrent.Stream(entry)
			return c.Stream(http.StatusOK, mime, stream)
		}
	})

	e.File("/movie/static", storage.Movie())
}
