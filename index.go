package main

import (
	"videoteque/cmd"
	"videoteque/server"
	"videoteque/subtitle"
)

func main() {
	cmd.Execute()
	subtitle.InitOpenSubtitlesIntegration()
	if server.Port != 0 {
		server.Init()
	}
}
