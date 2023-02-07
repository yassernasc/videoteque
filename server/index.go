package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/olahol/melody"
)

func Init() {
	e := echo.New()
	m := melody.New()

	e.HidePort = true
	e.HideBanner = true

	UiRoutes(e)
	MovieRoutes(e)
	SubtitleRoutes(e)
	WsRoutes(e, m)

	port := ":1313"
	fmt.Printf("lugosi is awake at http://localhost%v", port)
	e.Logger.Fatal(e.Start(port))
}
