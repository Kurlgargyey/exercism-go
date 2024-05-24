// Package isogram contains functionality for determining whether a string is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram takes a string and returns a bool indicating whether the string is an isogram.
func IsIsogram(word string) bool {
	freqs := make(map[rune]int)
	for _, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) {
			freqs[r]++
		}
	}
	return !any_value(freqs, func(count int) bool { return count > 1 })
}

func any_value[T interface{}, K comparable](mapping map[K]T, pred func(T) bool) bool {
	for _, v := range mapping {
		if pred(v) {
			return true
		}
	}
	return false
}
