package helpers

import (
	"strings"
	"unicode"
)

//spaceStringsBuilder delete all space in str
func SpaceStringsBuilder(str string) (string, error) {
	var b strings.Builder
	b.Grow(len(str))
	freeze := false
	for _, ch := range str {
		if ch == '"' {
			switch freeze {
			case true:
				freeze = false
			case false:
				freeze = true
			}
			b.WriteRune(ch)
			continue
		}
		if !freeze {
			if !unicode.IsSpace(ch) {
				b.WriteRune(ch)
			}
		} else {
			b.WriteRune(ch)
		}
	}
	return b.String(), nil
}
