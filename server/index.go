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

func showMessage() {
	ip := net.LocalIp()
	url := fmt.Sprintf("http://%v:%v", ip, storage.Port)
	fmt.Println("url:", url)

	if storage.ShowQrCode {
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

	p := fmt.Sprintf(":%v", storage.Port)
	e.Logger.Fatal(e.Start(p))
}
