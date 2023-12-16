package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mdp/qrterminal/v3"
	"github.com/olahol/melody"
	"lugosi/net"
	"os"
)

const port = ":1313"

var ip = net.LocalIp()
var url = "http://" + ip + port

func showMessage() {
	fmt.Printf("lugosi is awake at %v\n\n", url)
	fmt.Print("scan to open the settings page\n")
	qrterminal.Generate(url+"/settings", qrterminal.L, os.Stdout)
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
