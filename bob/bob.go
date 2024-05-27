package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	switch {
	case IsQuestion(remark) && IsYelling(remark):
		return "Calm down, I know what I'm doing!"
	case IsQuestion(remark):
		return "Sure."
	case IsYelling(remark):
		return "Whoa, chill out!"
	case IsSilence(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func IsQuestion(remark string) bool {
	return strings.HasSuffix(strings.TrimSpace(remark), "?")
}

func IsYelling(remark string) bool {
	found_letter := false
	for _, r := range remark {
		if unicode.IsLetter(r) {
			found_letter = true
		}
		if unicode.IsLower(r) {
			return false
		}
	}
	return found_letter
}

func IsSilence(remark string) bool {
	for _, r := range remark {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
