package tools

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"AHCII 7-bit chars":     {input: "aBcDeF12", want: "21FeDcBa"},
		"Russian chars":         {input: "РазРаботкА", want: "АктобаРзаР"},
		"Separators":            {input: ".,/ ", want: " /,."},
		"Chinese chars":         {input: "  中华人民共和国 ", want: " 国和共民人华中  "},
		"Special Unicode chars": {input: "✠⚛✪⚡", want: "⚡✪⚛✠"},
		"Diacritical chars":     {input: "ỗỆἣ", want: "ἣỆỗ"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Reverse(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
