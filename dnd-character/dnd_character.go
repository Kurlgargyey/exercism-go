package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score)-10)/2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	lowest := 6
	sum:= 0
	for i:= 0; i < 4; i++ {
		roll := rand.Intn(6)+1
		if lowest > roll {lowest = roll}
		sum += roll
	}
	return sum - lowest
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	consti := Ability()
	consti_mod := Modifier(consti)
	return Character{Ability(), Ability(), consti, Ability(), Ability(), Ability(), 10+consti_mod}
}
