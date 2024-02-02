package movie

import (
	"io"
	"os"
)

type localVideo struct {
	path string
}

func newLocalVideo(path string) Video {
	return localVideo{path}
}

func (l localVideo) Reader() io.ReadSeekCloser {
	f := l.file()
	return f
}

func (l localVideo) Path() string {
	return l.path
}

func (l localVideo) Size() int64 {
	f := l.file()
	defer f.Close()

	fi, _ := f.Stat()
	return fi.Size()
}

func (l localVideo) file() *os.File {
	f, _ := os.Open(l.path)
	return f
}
