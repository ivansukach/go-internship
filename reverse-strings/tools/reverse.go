/*
Package tools implements a simple library, that contains tools for work with strings.
*/
package tools

// Reverse is a function that reverses a string.
func Reverse(in string) string {
	tmp := make([]rune, len(in))
	amt := 0
	for _, v := range in {
		tmp[amt] = v
		amt++
	}
	tmp = tmp[:amt]
	var opp int
	for i := amt/2 - 1; i >= 0; i-- {
		opp = amt - 1 - i
		tmp[i], tmp[opp] = tmp[opp], tmp[i]
	}
	return string(tmp)
}
