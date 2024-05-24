// Package luhn provides functionality for validating numbers using the Luhn algorithm.
package luhn

import (
	"strconv"
	"unicode"
)

// Valid() takes a number represented by a string and returns its Luhn algorithm validation result.
func Valid(id string) bool {
	length := 0
	sum := 0
	runes := []rune(id)
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		switch {
		default:
			continue
		case !unicode.IsDigit(r) && !unicode.IsSpace(r):
			return false
		case unicode.IsDigit(r):
			length++
			num, _ := strconv.Atoi(string(r))
			if length%2 == 0 {
				num = num * 2
				if num > 9 {
					num = num - 9
				}
				sum += num
			} else {
				sum += num
			}
		}
	}
	if length <= 1 {
		return false
	}
	return sum%10 == 0
}
