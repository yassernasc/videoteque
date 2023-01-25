package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var dist embed.FS

func Content() (fs.FS, error) {
	return fs.Sub(dist, "dist")
}
