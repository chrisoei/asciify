package asciify

import (
	"code.google.com/p/go.text/unicode/norm"
	"regexp"
	"strings"
)

var isAscii = regexp.MustCompile("\\A[ -~]*\\z")
var nonAscii = regexp.MustCompile("[^ -~]+")
var replacer = strings.NewReplacer(
	"\u0060", "'",
	"\u00B4", "'",
	"\u2018", "'",
	"\u2019", "'",
	"\u201C", "\"",
	"\u201D", "\"")

func Convert(s string) string {
	return nonAscii.ReplaceAllLiteralString(
		replacer.Replace(norm.NFD.String(s)),
		"")
}

func IsAscii(s string) bool {
	return isAscii.MatchString(s)
}
