package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mdp/qrterminal/v3"
	"github.com/olahol/melody"
	"os"
	"videoteque/net"
	"videoteque/storage"
)

const port = ":1313"

func showMessage() {
	ip := net.LocalIp()
	url := "http://" + ip + port
	fmt.Println("url:", url)

	if storage.ShowQrCode() {
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
	e.Logger.Fatal(e.Start(port))
}
