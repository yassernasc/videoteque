package storage

import (
	"lugosi/judgment"
	"lugosi/net"
)

var movie string
var subtitle string

func SetMovie(entry string) {
	if judgment.IsUrl(entry) {
		entry = net.AvoidLocalhostNotation(entry)
	}

	movie = entry
}

func Movie() string {
	return movie
}

func SetSubtitle(entry string) {
	subtitle = entry
}

func Subtitle() string {
	return subtitle
}
