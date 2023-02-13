package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/olahol/melody"
	"lugosi/net"
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

	ip := net.LocalIp()
	if ip == "" {
		ip = "localhost"
	}

	port := ":1313"
	fmt.Printf("lugosi is awake at http://%v%v", ip, port)
	e.Logger.Fatal(e.Start(port))
}
