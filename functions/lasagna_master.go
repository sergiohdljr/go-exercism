package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepPerLayer int) int {
	if prepPerLayer == 0 {
		prepPerLayer = 2
	}
	numberOfLayers := len(layers)
	return numberOfLayers * prepPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(slice []string) (int, float64) {
	panic("error")
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	panic("error")
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	panic("error")
}
