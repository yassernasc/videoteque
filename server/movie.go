package server

import (
	"github.com/labstack/echo/v4"
	"lugosi/storage"
	"net/http"
)

func MovieRoutes(e *echo.Echo) {
	e.GET("/movie", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, storage.Movie)
	})
}
