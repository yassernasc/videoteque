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
		redirect := storage.Movie()
		if judgment.IsFile(redirect) {
			redirect = "/movie/static"
		}

		return c.Redirect(http.StatusMovedPermanently, redirect)
	})

	e.File("/movie/static", storage.Movie())
}
