package subtitle

import (
	"bytes"
	"github.com/asticode/go-astisub"
	iconv "github.com/djimenez/iconv-go"
	"github.com/gogs/chardet"
	"net/http"
	"os"
	"videoteque/fs"
)

var Entry string

func Get() (subtitle string, statusCode int, err error) {
	if shouldDownload() {
		err := downloadAutomatically()
		if err != nil {
			return "", http.StatusInternalServerError, err
		}
	}

	if Entry == "" {
		return "", http.StatusNoContent, nil
	}

	vtt := ensureVTT(Entry)
	vttUtf8 := ensureUTF8(vtt)

	return vttUtf8, http.StatusOK, nil
}

func IsValidEntry(path string) bool {
	pathExt := fs.Ext(path)
	validExtensions := [...]string{".srt", ".vtt"}

	for _, ext := range validExtensions {
		if ext == pathExt {
			return true
		}
	}

	return false
}

func ensureVTT(path string) string {
	if isSrt(path) {
		srt, _ := astisub.OpenFile(path)
		vttBuffer := &bytes.Buffer{}
		srt.WriteToWebVTT(vttBuffer)
		return vttBuffer.String()
	} else {
		vtt, _ := os.ReadFile(path)
		return string(vtt)
	}
}

func ensureUTF8(subtitle string) string {
	var isUTF8, charset = isUTF8(subtitle)

	if !isUTF8 {
		utf8, err := iconv.ConvertString(subtitle, charset, "utf-8")
		if err == nil {
			return utf8
		}
	}

	return subtitle
}

func isUTF8(data string) (bool, string) {
	detector := chardet.NewTextDetector()
	result, _ := detector.DetectBest([]byte(data))
	charset := result.Charset
	return charset == "UTF-8", charset
}

func isSrt(path string) bool {
	return fs.Ext(path) == ".srt"
}
