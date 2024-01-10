package movie

import (
	"errors"
	"github.com/cyruzin/golang-tmdb"
	ptn "github.com/middelink/go-parse-torrent-name"
	"lugosi/fs"
	"lugosi/torrent"
	"strconv"
)

var tc *tmdb.Client

func init() {
	key := "ac92176abc89a80e6f5df9510e326601"
	c, _ := tmdb.Init(key)
	c.SetClientAutoRetry()
	tc = c
}

type Type int
type Metadata struct {
	Backdrop string
	Episode  int
	Season   int
	Title    string
	Type     Type
	Year     int
}

const (
	Unknown Type = iota
	Movie
	Episode
)

func (e *Entry) LoadMetadata() {
	var filename string

	switch e.Format {
	case Magnet:
		path := torrent.InitClient(e.Payload)
		filename = fs.Filename(path)
	case File:
		filename = fs.Filename(e.Payload)
	}

	if filename != "" {
		metadata, err := infer(filename)
		if err == nil {
			e.Metadata = metadata
		}
	}
}

func infer(filename string) (*Metadata, error) {
	parsed, err := ptn.Parse(filename)
	if err != nil {
		return nil, err
	}

	m, err := populateParsed(parsed)
	if err != nil {
		return nil, err
	}

	getBackdrop(m)

	return m, nil
}

func populateParsed(p *ptn.TorrentInfo) (*Metadata, error) {
	var m Metadata
	m.Title = p.Title

	if p.Season != 0 && p.Episode != 0 {
		m.Season = p.Season
		m.Episode = p.Episode
		m.Type = Episode
		return &m, nil
	}

	if p.Year != 0 {
		m.Year = p.Year
		m.Type = Movie
		return &m, nil
	}

	return nil, errors.New("Infered data is incomplete")
}

func getBackdrop(m *Metadata) {
	var b string
	empty := make(map[string]string)

	switch m.Type {
	case Movie:
		options := map[string]string{"year": strconv.Itoa(m.Year)}
		res, err := tc.GetSearchMovies(m.Title, options)
		if err != nil && res.TotalResults == 0 {
			return
		}

		b = "w1280" + "/" + res.SearchMoviesResults.Results[0].BackdropPath
	case Episode:
		res, err := tc.GetSearchTVShow(m.Title, empty)
		if err != nil || res.TotalResults == 0 {
			return
		}

		showId := res.SearchTVShowsResults.Results[0].ID
		episode, err := tc.GetTVEpisodeDetails(int(showId), m.Season, m.Episode, empty)
		if err != nil {
			return
		}

		b = "original" + "/" + episode.StillPath

	}

	if b != "" {
		m.Backdrop = "https://www.themoviedb.org/t/p/" + b
	}
}
