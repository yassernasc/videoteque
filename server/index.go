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
	http.HandleFunc("/video", videoHandler)
	http.HandleFunc("/metadata", metadataHandler)
	http.HandleFunc("/subtitle", subtitleHandler)
	http.HandleFunc("/ws", handleWs)

	http.Handle("/", uiHandler())

	// page aliases, how to do automagically?
	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/settings.html", http.StatusMovedPermanently)
	})

	http.HandleFunc("/legacy", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/legacy.html", http.StatusMovedPermanently)
	})

	p := net.FormatPort(Port)
	panic(http.ListenAndServe(p, nil))
}
