package slices

import "slices"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2,6,9}
    return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
// ***TODO*** 
func GetItem(slice []int, index int) int {

     itemIndex := slices.Index(slice,index) 

     if itemIndex != -1 {
        return slice[itemIndex]
     }

    return itemIndex
}