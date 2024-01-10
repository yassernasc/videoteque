package server

import (
	"github.com/labstack/echo/v4"
	"lugosi/mime"
	"lugosi/movie"
	"lugosi/storage"
	"lugosi/torrent"
	"net/http"
)

func MovieRoutes(e *echo.Echo) {
	e.GET("/movie", func(c echo.Context) error {
		entry := storage.Movie()

		switch entry.Format {
		case movie.Magnet:
			stream, displayPath := torrent.Stream(entry.Payload)
			mime := mime.Get(displayPath)
			return c.Stream(http.StatusOK, mime, stream)
		case movie.File:
			return c.Redirect(http.StatusMovedPermanently, "/movie/static")
		default: // Url
			return c.Redirect(http.StatusMovedPermanently, entry.Payload)
		}
	})

	e.File("/movie/static", storage.Movie().Payload)

	e.GET("/metadata", func(c echo.Context) error {
		m := storage.Movie().Metadata

		if m == nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "No metadata to provide")
		}

		return c.JSON(http.StatusOK, m)
	})
}
