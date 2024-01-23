package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"videoteque/subtitle"
)

func SubtitleRoutes(e *echo.Echo) {
	e.GET("/subtitle", func(c echo.Context) error {
		s := subtitle.Get()
		if s == "" {
			return c.NoContent(http.StatusNotFound)
		}

		stream := strings.NewReader(s)
		return c.Stream(http.StatusOK, "text/vtt", stream)
	})
}
