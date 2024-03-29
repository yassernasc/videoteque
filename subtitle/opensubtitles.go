package subtitle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"videoteque/cache"
	"videoteque/fs"
	"videoteque/lang"
	"videoteque/movie"
)

const (
	agent         = "videoteque v0.1"
	apiKey        = "wWESRShMeQTiozKJrZlPX2nOlKsWiSpZ"
	baseUrl       = "https://api.opensubtitles.com/api/v1"
	tokenCacheKey = "os-token"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResult struct {
	Token string
}

type searchResult struct {
	TotalCount int `json:"total_count"`
	Data       []searchResultData
}

type searchResultData struct {
	Attributes searchResultDataAttributes
}

type searchResultDataAttributes struct {
	Language string
	Files    []searchResultDataAttributesFiles
}

type searchResultDataAttributesFiles struct {
	Id int `json:"file_id"`
}

type downloadArgs struct {
	Id     int    `json:"file_id"`
	Format string `json:"sub_format"`
}

type downloadResult struct {
	Link     string
	FileName string `json:"file_name"`
}

var credentials Credentials
var client = http.Client{}

func SetCredentials(username string, password string) {
	credentials = Credentials{username, password}
}

func shouldDownload() bool {
	return Entry == "" &&
		credentials.Username != "" &&
		credentials.Password != "" &&
		movie.MetadataRef != nil &&
		movie.MetadataRef.Tmdb != nil &&
		lang.UserLang != "" &&
		!lang.IsUserLang(movie.MetadataRef.Tmdb.Language)
}

func downloadAutomatically() error {
	t, err := login()
	if err != nil {
		return err
	}

	id, err := search()
	if err != nil {
		return err
	}

	return download(id, t)
}

func login() (token string, err error) {
	if t := cache.Read(tokenCacheKey); isJWTNotExpired(t) {
		return t, nil
	} else {
		cache.Delete(tokenCacheKey)
	}

	url := baseUrl + "/login"

	buf, _ := json.Marshal(credentials)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Api-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", agent)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("open subtitle login returned status %s", resp.Status)
	}
	defer resp.Body.Close()

	var result loginResult
	b, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(b, &result)
	if err != nil {
		return "", err
	}

	t := result.Token
	cache.Write(tokenCacheKey, []byte(t))

	return t, err
}

func search() (id int, err error) {
	url := fmt.Sprintf("%s/subtitles?%s", baseUrl, getSearchOptions())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("Api-Key", apiKey)
	req.Header.Add("User-Agent", agent)

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("open subtitle subtitle search returned status %s", resp.Status)
	}
	defer resp.Body.Close()
	var result searchResult
	b, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(b, &result)
	if err != nil {
		return 0, err
	}

	if result.TotalCount == 0 {
		return 0, errors.New("no subtitle found")
	}

	return result.Data[0].Attributes.Files[0].Id, nil
}

func getSearchOptions() string {
	meta := movie.MetadataRef

	v := url.Values{}
	v.Add("ai_translated", "exclude")
	v.Add("hearing_impaired", "exclude")
	v.Add("languages", lang.UserLang)
	v.Add("tmdb_id", strconv.Itoa(int(meta.Tmdb.Id)))

	if meta.Guessed.Type == movie.Movie {
		v.Add("type", "movie")
	} else {
		v.Add("type", "episode")
	}

	hash, err := moviehash()
	if err == nil {
		v.Add("moviehash", hash)
	}

	return v.Encode()
}

func download(subId int, token string) error {
	url := baseUrl + "/download"

	body := downloadArgs{subId, "srt"}
	buf, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Api-Key", apiKey)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", agent)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("open subtitle download returned status %s", resp.Status)
	}
	defer resp.Body.Close()

	var result downloadResult
	b, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(b, &result)
	if err != nil {
		return err
	}

	path, err := fs.DownloadTempFile(result.Link, result.FileName)
	if err != nil {
		return err
	}

	Entry = path
	return nil
}

func isJWTNotExpired(tokenString string) bool {
	if tokenString == "" {
		return false
	}

	t, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return false
	}

	exp, err := t.Claims.GetExpirationTime()
	if err != nil {
		return false
	}

	return time.Now().Before(exp.Time)
}
