package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	e.HidePort = true
	e.HideBanner = true

	UiRoutes(e)
	MovieRoutes(e)

	port := ":1313"
	fmt.Printf("lugosi is awake at http://localhost%v", port)
	e.Logger.Fatal(e.Start(port))
}
