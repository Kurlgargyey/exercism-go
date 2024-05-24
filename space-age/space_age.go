package space

type Planet string

var orbitals = map[Planet]float64{
	"Mercury": 0.240846,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const SecondsInYear = 365.25 * 24 * 60 * 60

func Age(seconds float64, planet Planet) float64 {
	orbital, exists := orbitals[planet]
	if !exists {
		return -1
	}
	return seconds / orbital / SecondsInYear
}
