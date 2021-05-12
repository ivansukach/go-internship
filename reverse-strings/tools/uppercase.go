package tools

import (
	"unicode"
)

// UpperCase converts string to upper case.
func UpperCase(in string) string {
	tmp := make([]rune, len(in))
	amt := 0
	for _, v := range in {
		tmp[amt] = unicode.ToUpper(v)
		amt++
	}
	return string(tmp[:amt])
}
