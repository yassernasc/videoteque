package movie

import (
	ptn "github.com/middelink/go-parse-torrent-name"
	"lugosi/fs"
	"lugosi/torrent"
)

type Metadata struct {
	Title   string
	Year    int
	Season  int
	Episode int
}

func (e Entry) LoadMetadata() {
	var filename string

	switch e.Format {
	case Magnet:
		path := torrent.InitClient(e.Payload)
		filename = fs.Filename(path)
	case File:
		filename = fs.Filename(e.Payload)
	}

	if filename != "" {
		metadata, err := parse(filename)
		if err != nil {
			e.Metadata = metadata
		}
	}
}

func parse(filename string) (*Metadata, error) {
	parsed, err := ptn.Parse(filename)

	if err != nil {
		return nil, err
	}

	m := &Metadata{parsed.Title, parsed.Year, parsed.Season, parsed.Episode}
	return m, nil
}
