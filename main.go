package main

import (
	"fmt"
	"github.com/sergiohdljr/go-exercism.git/slices"
)

func main() {
	cards := slices.RemoveItem([]int{3, 2, 6, 4, 8}, 2)
    fmt.Println(cards, "// Output: [3 2 4 8]")

}