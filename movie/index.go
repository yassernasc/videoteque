package movie

import (
	"errors"
	"lugosi/fs"
	"lugosi/net"
	"lugosi/torrent"
)

type Format int
type Entry struct {
	Payload  string
	Format   Format
	Metadata *Metadata
}

const (
	Magnet Format = iota
	File
	Url
)

func New(p string) *Entry {
	var e Entry

	f, _ := GetFormat(p)

	if f == Url {
		p = net.AvoidLocalhostNotation(p)
	}

	e.Payload = p
	e.Format = f
	e.LoadMetadata()

	return &e
}

func GetFormat(p string) (Format, error) {
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
