package server

import (
	"github.com/labstack/echo/v4"
	"lugosi/storage"
)

func SubtitleRoutes(e *echo.Echo) {
	e.File("/subtitle", storage.Subtitle)
}
