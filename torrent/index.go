package torrent

import (
	"github.com/anacrolix/torrent"
	"io"
	"videoteque/fs"
)

var movie *torrent.File

func InitClient(magnet string) string {
	c := createClient()

	t, _ := c.AddMagnet(magnet)
	<-t.GotInfo()

	movie = findMovie(t.Files())
	return movie.DisplayPath()
}

func Stream(magnet string) (io.Reader, string) {
	return movie.NewReader(), movie.DisplayPath()
}

func createClient() *torrent.Client {
	config := torrent.NewDefaultClientConfig()
	config.DataDir = fs.TempDir()

	c, _ := torrent.NewClient(config)
	return c
}

func findMovie(files []*torrent.File) *torrent.File {
	var m *torrent.File
	var maxSize int64

	for _, f := range files {
		if f.Length() > maxSize {
			m = f
			maxSize = f.Length()
		}
	}

	return m
}
