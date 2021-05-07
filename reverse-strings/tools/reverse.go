package tools

func Reverse(in string) string {
	tmp := make([]rune, len(in))
	amt := 0
	for _, v := range in {
		tmp[amt] = v
		amt++
	}
	tmp = tmp[:amt]
	reversed := make([]rune, amt)
	for i, v := range tmp {
		reversed[amt-i-1] = v
	}
	return string(reversed)
}
