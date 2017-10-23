package kata

import (
	"strings"
	"unicode/utf8"
)

// Accum desc: This time no story, no theory. The examples below show you how to write function accum:
// Examples:
// accum("abcd") // -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") // -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") // -> "C-Ww-Aaa-Tttt"
// The parameter of accum is a string which includes only letters from a..z and A..Z.
func Accum(s string) string {
	acc := make([]string, utf8.RuneCount([]byte(s)))
	for i, r := range s {
		n := strings.ToUpper(string(r))
		if i > 0 {
			n += strings.Repeat(strings.ToLower(string(r)), i)
		}
		acc[i] = n
	}
	return strings.Join(acc, "-")
}
