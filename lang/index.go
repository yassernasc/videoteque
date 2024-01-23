package lang

import "strings"

var UserLang string

func IsUserLang(l string) bool {
	if UserLang == "" {
		return false
	}
	return countryCode(UserLang) == countryCode(l)
}

func isEqual(lang1 string, lang2 string) bool {
	return countryCode(lang1) == countryCode(lang2)
}

func countryCode(lang string) string {
	return strings.ToLower(strings.Split(lang, "-")[0])
}
