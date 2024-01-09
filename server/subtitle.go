package server

import (
	"github.com/labstack/echo/v4"
	"lugosi/storage"
	"lugosi/subtitle"
	"net/http"
	"strings"
)

func SubtitleRoutes(e *echo.Echo) {
	e.GET("/subtitle", func(c echo.Context) error {
		path := storage.Subtitle()

		if path == "" {
			return c.NoContent(http.StatusNotFound)
		}

		s := subtitle.Get(path)
		stream := strings.NewReader(s)
		return c.Stream(http.StatusOK, "text/vtt", stream)
	})
}
