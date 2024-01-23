package movie

import (
	"fmt"
	"github.com/cyruzin/golang-tmdb"
	"strconv"
	"videoteque/lang"
)

type tmdbInfo struct {
	Backdrop string
	Id       int64
	Language string
	Overview string
	Title    string
}

var tc *tmdb.Client

func init() {
	key := "ac92176abc89a80e6f5df9510e326601"
	c, _ := tmdb.Init(key)
	c.SetClientAutoRetry()
	tc = c
}

func fetchTmdbInfo(m *guessed) (*tmdbInfo, error) {
	if m.Type == Episode {
		return searchEpisode(m.Title, m.Season, m.Episode)
	}

	return searchMovie(m.Title, m.Year)
}

func searchMovie(title string, year int) (*tmdbInfo, error) {
	options := map[string]string{"year": strconv.Itoa(year)}

	res, err := tc.GetSearchMovies(title, options)
	if err != nil {
		return nil, err
	}

	if res.TotalResults == 0 {
		return nil, fmt.Errorf("tmdb cannot find the movie %s", title)
	}

	m := res.SearchMoviesResults.Results[0]

	b := tmdb.GetImageURL(m.BackdropPath, tmdb.Original)
	t := getTitle(m.OriginalLanguage, m.OriginalTitle, m.Title)
	i := tmdbInfo{b, m.ID, m.OriginalLanguage, m.Overview, t}

	return &i, nil
}

func searchEpisode(title string, season int, episode int) (*tmdbInfo, error) {
	params := make(map[string]string)

	res, err := tc.GetSearchTVShow(title, params)
	if err != nil {
		return nil, err
	}

	if res.TotalResults == 0 {
		return nil, fmt.Errorf("tmdb cannot find the tv show %s", title)
	}

	t := res.SearchTVShowsResults.Results[0]
	tId := int(t.ID)

	e, err := tc.GetTVEpisodeDetails(tId, season, episode, params)
	if err != nil {
		return nil, err
	}

	b := tmdb.GetImageURL(e.StillPath, tmdb.Original)
	i := tmdbInfo{b, e.ID, t.OriginalLanguage, e.Overview, e.Name}

	return &i, nil
}

func getTitle(l string, originalTitle string, defaultTitle string) string {
	if lang.IsUserLang(l) {
		return originalTitle
	}

	return defaultTitle
}
