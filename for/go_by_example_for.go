package forGo

import "fmt"

func Go_by_example_for(){
	fmt.Println("The most basic type, with a single condition.")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("A classic initial/condition/after for loop.")
    for j := 0 ; j < 3 ; j++ {
		fmt.Println(j+1)
	}

	
}