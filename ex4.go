package main

import "fmt"

// Note the array length must be specified
func double(arr [3]int) [3]int {
	for index, value := range arr {
		fmt.Println(index, ": ", value)
		arr[index] = value * 2
	}
	return arr
}

func main() {
	a := [...]int{1, 2, 3}
	b := a   // b is a duplicate of a
	a[0] = 4 // this modifies a[] and not b[]
	fmt.Println("len(a): ", len(a))
	fmt.Println("\nA: ", a)
	fmt.Println("B: ", b)

	c := double(b)
	fmt.Println("\nC: ", c) // New array returned
	fmt.Println("B:", b)    // Original array b
}
