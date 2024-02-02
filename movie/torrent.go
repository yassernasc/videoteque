package movie

import (
	anacrolix "github.com/anacrolix/torrent"
	"io"
	"videoteque/torrent"
)

type torrentVideo struct {
	file *anacrolix.File
}

func newTorrentVideo(entry string) Video {
	f := torrent.AddMagnet(entry)
	return torrentVideo{f}
}

func (t torrentVideo) Reader() io.ReadSeekCloser {
	return t.file.NewReader()
}

func (t torrentVideo) Path() string {
	return t.file.DisplayPath()
}

func (t torrentVideo) Size() int64 {
	return t.file.Length()
}
