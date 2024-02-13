package fs

import (
	"github.com/adrg/xdg"
	"path/filepath"
)

const appFolder = "videoteque"

func ConfigDir() string {
	return filepath.Join(xdg.ConfigHome, appFolder)
}

func CacheDir() string {
	return filepath.Join(xdg.CacheHome, appFolder)
}
