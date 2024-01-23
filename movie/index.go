package movie

import (
	"errors"
	"videoteque/fs"
	"videoteque/net"
	"videoteque/torrent"
)

type format int
type video struct {
	Payload  string
	Format   format
	Metadata *metadata
}

const (
	Magnet format = iota
	File
	Url
)

var Video video

func Init(entry string) {
	var v video

	f, _ := GetFormat(entry)

	if f == Url {
		entry = net.AvoidLocalhostNotation(entry)
	}

	v.Payload = entry
	v.Format = f
	v.loadMetadata()

	Video = v
}

func GetFormat(p string) (format, error) {
	switch {
	case torrent.IsMagnetLink(p):
		return Magnet, nil
	case fs.IsFile(p):
		return File, nil
	case net.IsUrl(p):
		return Url, nil
	default:
		return Url, errors.New("invalid movie entry format")
	}
}
