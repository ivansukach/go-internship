package tools

import (
	"testing"
)

// TestReverse tests the string reverse method implementation in a multiple cases.
func TestReverse(t *testing.T) {
	tests := map[string]struct {
		input  string
		expRes string
	}{
		"AHCII 7-bit chars":     {input: "aBcDeF12", expRes: "21FeDcBa"},
		"Russian chars":         {input: "РазРаботкА", expRes: "АктобаРзаР"},
		"Separators":            {input: ".,/ ", expRes: " /,."},
		"Chinese chars":         {input: "  中华人民共和国 ", expRes: " 国和共民人华中  "},
		"Special Unicode chars": {input: "✠⚛✪⚡", expRes: "⚡✪⚛✠"},
		"Diacritical chars":     {input: "ỗỆἣ", expRes: "ἣỆỗ"},
		"Empty string":          {input: "", expRes: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := Reverse(tc.input)
			if tc.expRes != result {
				t.Errorf("Reverse(%s) = %s; expRes %s", tc.input, result, tc.expRes)
			}
		})
	}
}
