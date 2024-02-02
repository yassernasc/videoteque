package movie

import (
	"errors"
	"io"
	"videoteque/fs"
	"videoteque/torrent"
)

type Video interface {
	Reader() io.ReadSeekCloser
	Path() string
	Size() int64
}

var VideoRef Video

func Init(entry string) {
	v, _ := NewVideo(entry)
	loadMetadata(v)
	VideoRef = v
}

func NewVideo(entry string) (Video, error) {
	switch {
	case torrent.IsMagnetLink(entry):
		return newTorrentVideo(entry), nil
	case fs.IsFile(entry):
		return newLocalVideo(entry), nil
	default:
		return nil, errors.New("invalid movie entry format")
	}
}
