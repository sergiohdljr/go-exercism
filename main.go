package main

import (
	"fmt"

	lasagna "github.com/sergiohdljr/go-exercism.git/functions"
)

func main() {

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	fmt.Printf(lasagna.AddSecretIngredient(friendsList, myList))

}
