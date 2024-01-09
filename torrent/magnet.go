package torrent

import "regexp"

func IsMagnetLink(str string) bool {
	pattern := `(?i)magnet:\?xt=urn:[a-z0-9]+:[a-z0-9]{32}`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
