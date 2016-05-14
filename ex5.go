package main

import "fmt"

func double(arr []int) []int {
	for index, value := range arr {
		fmt.Println(index, ": ", value)
		arr[index] = value * 2
	}
	return arr
}

func main() {
	a := make([]int, 3, 10)
	var b [10]int

	fmt.Println("Length of A: ", len(a))
	fmt.Println("Capacity of A: ", cap(a))
	fmt.Println("A: ", a)

	b = [10]int{1, 2, 3, 4, 5}
	fmt.Println("\nLength of B: ", len(b))
	fmt.Println("Capacity of B: ", cap(b))
	fmt.Println("B: ", b)

	c := b[0:5]
	fmt.Println("\nLength of C: ", len(c))
	fmt.Println("Capacity of C: ", cap(c))
	fmt.Println("C: ", c)

	// Assigning a slice does not copy the values, d and c point to the same underlying array
	d := c
	fmt.Println("\nLength of D: ", len(d))
	fmt.Println("Capacity of D: ", cap(d))
	fmt.Printf("Address of C: %p and D: %p\n", &c, &d) // C and D are distinct variables
	fmt.Println("However, they point the the same place in memory:")
	fmt.Printf("C (%p): %x\n", c, c)
	fmt.Printf("D (%p): %x\n", d, d)

	fmt.Println("\nAssign 100 to d[0] and we get:")
	d[0] = 100
	fmt.Println("C: ", c)
	fmt.Println("D: ", d)

	fmt.Println("Make D a new slice, twice the capacity of C and copy C into D")
	d = make([]int, len(c), 2*cap(c))
	copy(d, c) // target, source ; Read like as assignment, d = c
	fmt.Println("\nLength of D: ", len(d))
	fmt.Println("Capacity of D: ", cap(d))
	fmt.Println("D: ", d)

	fmt.Println("\nAssign 10 to d[0] and we get:")
	d[0] = 10
	fmt.Println("C: ", c)
	fmt.Println("D: ", d)

	fmt.Println("\nLet's double D")
	fmt.Println(double(d))

}
