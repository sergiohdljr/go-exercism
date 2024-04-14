package main

import (
	"fmt"

	lasagna "github.com/sergiohdljr/go-exercism.git/functions"
)

func main() {

	quantities := []float64{0.5, 250, 150, 3, 0.5}
	scaledQuantities := lasagna.ScaleRecipe(quantities, 6)
	fmt.Println(scaledQuantities)
	// => []float64{1.5, 750, 450, 9, 1.5}

}
