package main

import (
	"fmt"

	lasagna "github.com/sergiohdljr/go-exercism.git/functions"
)

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}

	fmt.Println(lasagna.PreparationTime(layers, 3))

	fmt.Println(lasagna.PreparationTime(layers, 0))

}
