package judgment

import (
	"github.com/gogs/chardet"
	"net/url"
	"os"
	"path/filepath"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsFile(str string) bool {
	_, err := os.Stat(str)
	return err == nil
}

func IsSubtitle(path string) bool {
	pathExt := filepath.Ext(path)
	validExtensions := [...]string{".srt", ".vtt"}

	for _, ext := range validExtensions {
		if ext == pathExt {
			return true
		}
	}

	return false
}

func IsSrt(path string) bool {
	return filepath.Ext(path) == ".srt"
}

func IsUTF8(data string) (bool, string) {
	detector := chardet.NewTextDetector()
	result, _ := detector.DetectBest([]byte(data))
	charset := result.Charset
	return charset == "UTF-8", charset
}
