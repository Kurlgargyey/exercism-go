package scrabble

import "strings"

func Score(word string) int {
	score := 0
	for _, r := range strings.ToUpper(word) {
		score += value(r)
	}
	return score
}

func value(r rune) int {
	switch r {
	default:
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
}
