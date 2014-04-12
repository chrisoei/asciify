package asciify

import (
	"regexp"
	"strings"
)

func Convert(s string) string {
	return Replacer().Replace(s)
}

func IsAscii(s string) bool {
	return regexp.MustCompile("\\A[ -~]*\\z").MatchString(s)
}

func Replacer() *strings.Replacer {
	return strings.NewReplacer(
		"\u0060", "'",
		"\u00B4", "'",
		"\u2018", "'",
		"\u2019", "'",
		"\u201C", "\"",
		"\u201D", "\"")
}
