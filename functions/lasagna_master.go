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
	secretIncredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIncredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	copyQuantities := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		copyQuantities[i] = quantities[i]  / 2 * float64(portions)
	}

	return copyQuantities
}
