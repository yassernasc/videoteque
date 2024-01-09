package storage

import "lugosi/movie"

var entry *movie.Entry
var subtitle string
var showQrCode bool

func SetMovie(e *movie.Entry) {
	entry = e
}

func Movie() *movie.Entry {
	return entry
}

func SetSubtitle(entry string) {
	subtitle = entry
}

func Subtitle() string {
	return subtitle
}

func SetShowQrCode(s bool) {
	showQrCode = s
}

func ShowQrCode() bool {
	return showQrCode
}
