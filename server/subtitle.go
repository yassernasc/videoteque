package server

import (
	"bytes"
	"github.com/asticode/go-astisub"
	iconv "github.com/djimenez/iconv-go"
	"github.com/labstack/echo/v4"
	"lugosi/judgment"
	"lugosi/storage"
	"net/http"
	"os"
	"strings"
)

func SubtitleRoutes(e *echo.Echo) {
	e.GET("/subtitle", func(c echo.Context) error {
		path := storage.Subtitle()

		if path == "" {
			return c.NoContent(http.StatusNotFound)
		}

		vtt := ensureVTT(path)
		vttUtf8 := ensureUTF8(vtt)

		stream := strings.NewReader(vttUtf8)
		return c.Stream(http.StatusOK, "text/vtt", stream)
	})
}

func ensureVTT(path string) string {
	if judgment.IsSrt(path) {
		srt, _ := astisub.OpenFile(path)
		vttBuffer := &bytes.Buffer{}
		srt.WriteToWebVTT(vttBuffer)
		return vttBuffer.String()
	} else {
		vtt, _ := os.ReadFile(path)
		return string(vtt)
	}
}

func ensureUTF8(subtitle string) string {
	var isUTF8, charset = judgment.IsUTF8(subtitle)

	if !isUTF8 {
		utf8, err := iconv.ConvertString(subtitle, charset, "utf-8")
		if err == nil {
			return utf8
		}
	}

	return subtitle
}
