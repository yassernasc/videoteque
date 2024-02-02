package movie

import (
	"errors"
	ptn "github.com/middelink/go-parse-torrent-name"
	"videoteque/fs"
)

type kind int

type guessed struct {
	Episode int
	Season  int
	Title   string
	Type    kind
	Year    int
}

type Metadata struct {
	Guessed *guessed
	Tmdb    *tmdbInfo
}

const (
	Movie kind = iota
	Episode
)

var MetadataRef *Metadata

func loadMetadata(v Video) {
	filename := fs.Filename(v.Path())
	if m, err := infer(filename); err == nil {
		MetadataRef = m
	}
}

func infer(filename string) (*Metadata, error) {
	var m Metadata

	parsed, err := ptn.Parse(filename)
	if err != nil {
		return nil, err
	}

	g, err := populateParsed(parsed)
	if err != nil {
		return nil, err
	} else {
		m.Guessed = g
	}

	t, err := fetchTmdbInfo(g)
	if err != nil {
		return nil, err
	} else {
		m.Tmdb = t
	}

	return &m, nil
}

func populateParsed(p *ptn.TorrentInfo) (*guessed, error) {
	var g guessed
	g.Title = p.Title

	if p.Season != 0 && p.Episode != 0 {
		g.Season = p.Season
		g.Episode = p.Episode
		g.Type = Episode
		return &g, nil
	}

	if p.Year != 0 {
		g.Year = p.Year
		g.Type = Movie
		return &g, nil
	}

	return nil, errors.New("infered data is incomplete")
}
