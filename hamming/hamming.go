// Package hamming allows the user to calculate the hamming distance between two DNA sequences presented as strings.
package hamming

import "errors"

//Distance takes two strings and returns their hamming distance. If the strings are of different lengths, an error is returned. In the error case, the function returns 0 for the hamming distance.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequence lengths do not match")
	}
	distance := 0
	for i, r := range a {
		if []rune(b)[i] != r {
			distance++
		}
	}
	return distance, nil
}
