package asciify

import (
    "regexp"
)

func Convert(s string) string {
	return s
}

func IsAscii(s string) bool {
    return regexp.MustCompile("\\A[A-Za-z0-9 ]*\\z").MatchString(s)
}
