package romannumerals

import (
	"errors"
	"math"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 {
		return "", errors.New("only positive numbers are allowed")
	}
	if input >= 4000 {
		return "", errors.New("only numbers below 4000 are allowed")
	}

	result, thousands_remainder := RomanizeThousands(input)

	hundreds_string, hundreds_remainder := RomanizeHundreds(thousands_remainder)
	result += hundreds_string

	tens_string, tens_remainder := RomanizeTens(hundreds_remainder)
	result += tens_string

	result += RomanizeOnes(tens_remainder)

	return result, nil
}

func RomanizeThousands(input int) (string, int) {
	place := 1000.0
	place_val := math.Floor(float64(input) / place)
	remainder := input % int(place)

	return strings.Repeat("M", int(place_val)), remainder
}

func RomanizeHundreds(input int) (string, int) {
	place := 100.0
	place_val := math.Floor(float64(input) / place)
	remainder := input % int(place)

	switch {
	case place_val <= 3:
		return strings.Repeat("C", int(place_val)), remainder
	case place_val == 4:
		return "CD", remainder
	case place_val == 5:
		return "D", remainder
	case place_val == 9:
		return "CM", remainder
	default:
		return "D" + strings.Repeat("C", int(place_val)%5), remainder
	}

}

func RomanizeTens(input int) (string, int) {
	place := 10.0
	place_val := math.Floor(float64(input) / place)
	remainder := input % int(place)

	switch {
	case place_val <= 3:
		return strings.Repeat("X", int(place_val)), remainder
	case place_val == 4:
		return "XL", remainder
	case place_val == 5:
		return "L", remainder
	case place_val == 9:
		return "XC", remainder
	default:
		return "L" + strings.Repeat("X", int(place_val)%5), remainder
	}

}

func RomanizeOnes(input int) string {
	switch {
	case input <= 3:
		return strings.Repeat("I", input)
	case input == 4:
		return "IV"
	case input == 5:
		return "V"
	case input == 9:
		return "IX"
	default:
		return "V" + strings.Repeat("I", int(input)%5)
	}
}
