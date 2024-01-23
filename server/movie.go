package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"videoteque/fs"
	"videoteque/movie"
	"videoteque/torrent"
)

func MovieRoutes(e *echo.Echo) {
	e.GET("/movie", func(c echo.Context) error {
		entry := movie.Video

		switch entry.Format {
		case movie.Magnet:
			stream, displayPath := torrent.Stream(entry.Payload)
			mime := getMime(displayPath)
			return c.Stream(http.StatusOK, mime, stream)
		case movie.File:
			return c.Redirect(http.StatusMovedPermanently, "/movie/static")
		default: // Url
			return c.Redirect(http.StatusMovedPermanently, entry.Payload)
		}
	})

	e.File("/movie/static", movie.Video.Payload)

	e.GET("/metadata", func(c echo.Context) error {
		metadata := movie.Video.Metadata

		if metadata == nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "No metadata to provide")
		}

		return c.JSON(http.StatusOK, metadata)
	})
}

func getMime(filename string) string {
	// https://developer.mozilla.org/en-US/docs/Web/Media/Formats/Containers#browser_compatibility
	switch fs.Ext(filename) {
	case ".3gp":
		return "video/3gpp"
	case ".m4p", ".m4v", ".mp4":
		return "video/mp4"
	case ".mpeg", ".mpg":
		return "video/mpeg"
	case ".ogg", ".ogv":
		return "video/ogg"
	case ".webm":
		return "video/webm"
	default:
		return "video/mp4" // use mp4 mime as fallback
	}
}
