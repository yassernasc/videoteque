package server

import (
	"fmt"
	"github.com/mdp/qrterminal/v3"
	"net/http"
	"os"
	"videoteque/net"
)

var ShowQrCode bool
var Port int

func Init() {
	showMessage()
	startServer()
}

func showMessage() {
	ip := net.LocalIp()
	url := fmt.Sprintf("http://%v:%v", ip, Port)
	fmt.Println("url:", url)

	if ShowQrCode {
		fmt.Print("\n\nscan to open the settings page\n")
		qrterminal.Generate(url+"/settings", qrterminal.L, os.Stdout)
	}
}

func startServer() {
	http.HandleFunc("/", uiHandler)
	http.HandleFunc("/ws", handleWs)

	http.HandleFunc("/movie", videoHandler)
	http.HandleFunc("/metadata", metadataHandler)
	http.HandleFunc("/subtitle", subtitleHandler)

	p := net.FormatPort(Port)
	panic(http.ListenAndServe(p, nil))
}
