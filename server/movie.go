package server

import (
	"github.com/labstack/echo/v4"
	"lugosi/judgment"
	"lugosi/storage"
	"net/http"
)

func MovieRoutes(e *echo.Echo) {
	e.GET("/movie", func(c echo.Context) error {
		// check if the current movie is a file or a link
		redirect := "/static"
		if judgment.IsUrl(storage.Movie) {
			redirect = storage.Movie
		}

		return c.Redirect(http.StatusMovedPermanently, redirect)

	})

	e.File("/static", storage.Movie)
}
