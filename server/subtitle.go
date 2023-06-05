package server

import (
	"bytes"
	"github.com/asticode/go-astisub"
	"github.com/labstack/echo/v4"
	"lugosi/judgment"
	"lugosi/storage"
	"net/http"
)

func SubtitleRoutes(e *echo.Echo) {
	e.GET("/subtitle", func(c echo.Context) error {
		if storage.Subtitle() == "" {
			return c.NoContent(http.StatusNotFound)
		}

		if judgment.IsSrt(storage.Subtitle()) {
			// convert to vtt
			srt, _ := astisub.OpenFile(storage.Subtitle())
			var buf = &bytes.Buffer{}
			srt.WriteToWebVTT(buf)
			return c.Stream(http.StatusOK, "text/vtt", buf)
		}

		return c.Redirect(http.StatusMovedPermanently, "/subtitle/static")
	})

	e.File("/subtitle/static", storage.Subtitle())
}
