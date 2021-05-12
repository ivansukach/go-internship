/*
Package tools implements a simple library, that contains tools for work with strings.
*/
package tools

//Reverse is a function that reverses a string
func Reverse(in string) (result string) {
	for _, v := range in {
		result = string(v) + result
	}
	return
}
