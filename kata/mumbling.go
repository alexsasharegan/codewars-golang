package kata

import (
	"strings"
	"unicode/utf8"
)

// Accum kata.
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
