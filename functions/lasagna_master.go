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
	var noodlesQuantity int
	var sauce float64

	for i := 0; i < len(slice); i++ {
		if slice[i] == "noodles" {
			noodlesQuantity = noodlesQuantity + 1
		}
		if slice[i] == "sauce" {
			sauce = sauce + 1
		}
	}

	return noodlesQuantity * 50, sauce * 0.2

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	panic("error")
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	panic("error")
}
