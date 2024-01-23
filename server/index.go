package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mdp/qrterminal/v3"
	"github.com/olahol/melody"
	"os"
	"videoteque/net"
)

var ShowQrCode bool
var Port int

func showMessage() {
	ip := net.LocalIp()
	url := fmt.Sprintf("http://%v:%v", ip, Port)
	fmt.Println("url:", url)

	if ShowQrCode {
		fmt.Print("\n\nscan to open the settings page\n")
		qrterminal.Generate(url+"/settings", qrterminal.L, os.Stdout)
	}
}

func Init() {
	e := echo.New()
	m := melody.New()

	e.HidePort = true
	e.HideBanner = true

	UiRoutes(e)
	MovieRoutes(e)
	SubtitleRoutes(e)
	WsRoutes(e, m)

	showMessage()

	p := net.FormatPort(Port)
	e.Logger.Fatal(e.Start(p))
}
