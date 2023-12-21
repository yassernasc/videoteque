package torrent

import (
	"github.com/anacrolix/torrent"
	"io"
	"os"
	"strings"
)

func Stream(magnet string) (io.Reader, string) {
	c := createClient()

	t, _ := c.AddMagnet(magnet)
	<-t.GotInfo()

	m := findMovie(t.Files())
	return m.NewReader(), getMime(m.DisplayPath())
}

func createClient() *torrent.Client {
	config := torrent.NewDefaultClientConfig()
	config.DataDir = os.TempDir()

	c, _ := torrent.NewClient(config)
	return c
}

func findMovie(files []*torrent.File) *torrent.File {
	var movie *torrent.File
	var maxSize int64

	for _, f := range files {
		if f.Length() > maxSize {
			movie = f
			maxSize = f.Length()
		}
	}

	return movie
}

func getMime(filename string) string {
	// https://developer.mozilla.org/en-US/docs/Web/Media/Formats/Containers#browser_compatibility
	mimeMap := map[string]string{
		"3gp":  "video/3gpp",
		"m4p":  "video/mp4",
		"m4v":  "video/mp4",
		"mp4":  "video/mp4",
		"mpeg": "video/mpeg",
		"mpg":  "video/mpeg",
		"ogg":  "video/ogg",
		"ogv":  "video/ogg",
		"webm": "video/webm",
	}

	ext := getExt(filename)

	mime, ok := mimeMap[ext]
	if !ok {
		mime = mimeMap["mp4"] // use mp4 mime as fallback
	}

	return mime
}

func getExt(filename string) string {
	pieces := strings.Split(filename, ".")
	return pieces[len(pieces)-1]
}
