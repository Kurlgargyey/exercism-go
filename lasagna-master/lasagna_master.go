package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layer_time int) int {
	if layer_time > 0 {
		return len(layers) * layer_time
	}
	return len(layers) * 2
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		default:
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, ownList []string) {
	ownList[len(ownList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var scaled []float64 = make([]float64, len(amounts))
	base_portions := 2.0
	for i := 0; i < len(amounts); i++ {
		scaled[i] = amounts[i] * float64(portions) / base_portions
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
