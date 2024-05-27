package pangram

import (
	"strings"
	"unicode"
)

const all_letters int = 0b11111111111111111111111111

func IsPangram(input string) bool {
	letters := 0
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			letters |= 1 << (r-'a')
		}
	}
	return all_letters & letters == all_letters
}
