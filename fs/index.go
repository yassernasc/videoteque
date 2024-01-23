package fs

import (
	"github.com/adrg/xdg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const appFolder = "videoteque"

func ConfigDir() string {
	return filepath.Join(xdg.ConfigHome, appFolder)
}

func IsFile(str string) bool {
	_, err := os.Stat(str)
	return err == nil
}

func Filename(path string) string {
	b := filepath.Base(path)
	f := strings.TrimSuffix(b, Ext(b))
	return f
}

func Ext(filename string) string {
	return filepath.Ext(filename)
}

func TempDir() string {
	return os.TempDir()
}

func DownloadTempFile(url string, filename string) (path string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	p := filepath.Join(TempDir(), filename)
	out, err := os.Create(p)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return p, nil
}
