package judgment

import (
	"net/url"
	"os"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsFile(str string) bool {
	_, err := os.Stat(str)
	return err == nil
}
