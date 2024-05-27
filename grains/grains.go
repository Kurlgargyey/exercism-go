package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 {return 0, errors.New("no squares have no grains")}
	if number > 64 {return 0, errors.New("there are only 64 squares on the board")}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i:= 1; i <= 64; i++ {
		square, _ := Square(i)
		sum += square
	}
	return sum
}
