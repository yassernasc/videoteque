package torrent

import (
	"github.com/anacrolix/torrent"
	"videoteque/fs"
)

var client *torrent.Client
var video *torrent.File

func init() {
	config := torrent.NewDefaultClientConfig()
	config.DataDir = fs.TempDir()

	c, _ := torrent.NewClient(config)
	client = c
}

func AddMagnet(magnet string) *torrent.File {
	t, _ := client.AddMagnet(magnet)
	<-t.GotInfo()

	video = findVideo(t.Files())
	video.Download()
	return video
}

func findVideo(files []*torrent.File) *torrent.File {
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
