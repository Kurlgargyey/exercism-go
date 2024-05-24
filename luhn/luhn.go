package luhn

import (
	"strconv"
	"unicode"
)

func Valid(id string) bool {
	double_next := false
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
			if double_next {
				num = num * 2
				if num > 9 {
					num = num - 9
				}
				sum += num
				double_next = false
			} else {
				sum += num
				double_next = true
			}
		}
	}
	if length <= 1 {
		return false
	}
	return sum%10 == 0
}
