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
    if index < 0 || index >= len(slice){
        return -1
    }
    return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	
    if index < 0 || index >= len(slice){
        appendItem := append(slice,value)
        return appendItem
    }

    slice[index] = value

    return  slice

}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var valuesSlice []int
    valuesSlice = append(valuesSlice, values...)
    slice = append(valuesSlice, slice...)
    
    
    return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
 
    if index < 0 || index >= len(slice){
       return slice    
    }
    
	return slices.Delete(slice,index,(index+1))
}