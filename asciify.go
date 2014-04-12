package asciify

import (
	"regexp"
	"strings"
)

func Convert(s string) string {
	for k, v := range TranslationMap() {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}

func IsAscii(s string) bool {
	return regexp.MustCompile("\\A[ -~]*\\z").MatchString(s)
}

func TranslationMap() map[string]string {
	return map[string]string{
		"\u0060": "'",
		"\u00B4": "'",
		"\u2018": "'",
		"\u2019": "'",
		"\u201C": "\"",
		"\u201D": "\"",
	}
}
