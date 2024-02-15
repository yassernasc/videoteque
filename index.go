package main

import (
	"videoteque/cmd"
	"videoteque/server"
)

func main() {
	cmd.Execute()
	if server.Port != 0 {
		server.Init()
	}
}
