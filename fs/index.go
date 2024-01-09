package fs

import (
	"os"
	"path/filepath"
	"strings"
)

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
