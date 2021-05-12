package tools

import (
	"unicode"
)

// UpperCase converts string to upper case.
func UpperCase(in string) string {
	tmp := make([]rune, len(in))
	amt := 0
	for _, v := range in {
		tmp[amt] = v
		amt++
	}
	tmp = tmp[:amt]
	for i, v := range tmp {
		tmp[i] = unicode.ToUpper(v)
	}
	return string(tmp)
}
