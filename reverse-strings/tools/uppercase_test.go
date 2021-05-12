package tools

import (
	"fmt"
	"testing"
)

func TestUpperCase(t *testing.T) {
	tests := map[string]struct {
		input  string
		expRes string
	}{
		"AHCII 7-bit chars":     {input: "aBcDeF12", expRes: "ABCDEF12"},
		"Russian chars":         {input: "РазРаботкА", expRes: "РАЗРАБОТКА"},
		"Separators":            {input: ".,/ ", expRes: ".,/ "},
		"Chinese chars":         {input: "  中华人民共和国 ", expRes: "  中华人民共和国 "},
		"Special Unicode chars": {input: "✠⚛✪⚡", expRes: "✠⚛✪⚡"},
		"Empty string":          {input: "", expRes: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := UpperCase(tc.input)
			if tc.expRes != result {
				t.Errorf("UpperCase(%s) = %s; but exected %s", tc.input, result, tc.expRes)
			}
		})
	}
}

func ExampleUpperCase() {
	fmt.Println(UpperCase("hello"))
	// Output: HELLO
}
