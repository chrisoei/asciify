package asciify

import (
    "regexp"
)

func Convert(s string) string {
	return s
}

func IsAscii(s string) bool {
    return regexp.MustCompile("\\A[ -~]*\\z").MatchString(s)
}
