package main

import (
	"github.com/sergiohdljr/go-exercism.git/strings"
)

func main() {
	
	message := `
             **************************
             *    BUY NOW, SAVE 10%   *
             ************************** `

	print(strings.CleanupMessage(message))
}