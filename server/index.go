package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"lugosi/ui"
	"net/http"
)

func getUiContentHandler() http.Handler {
	content, err := ui.Content()
	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(http.FS(content))
}

func Init() {
	e := echo.New()

	e.HidePort = true
	e.HideBanner = true

	uiContentHandler := getUiContentHandler()
	e.GET("/", echo.WrapHandler(uiContentHandler))
	e.GET("/_next/*", echo.WrapHandler(getUiContentHandler()))

	port := ":1313"
	fmt.Printf("lugosi is awake at http://localhost%v", port)
	e.Logger.Fatal(e.Start(port))
}
