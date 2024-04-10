package lasagna

func PreparationTime(layers []string, prepPerLayer int) int {
	if prepPerLayer == 0 {
		prepPerLayer = 2
	}
	numberOfLayers := len(layers)
	return numberOfLayers * prepPerLayer
}
