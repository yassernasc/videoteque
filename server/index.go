package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	e.HidePort = true
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := ":1313"
	fmt.Printf("lugosi is awake at http://localhost%v", port)
	e.Logger.Fatal(e.Start(port))
}
