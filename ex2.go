package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func main() {
	var x int
	var a string = "A"
	var b = "B"
	y := 2 // Most often used in loops/conditionals/function return values

	x = 5 // What would happen if we comment this out??
	z := mult(x, y)
	fmt.Printf("%d * %d = %d\n", x, y, z)
	a = "foo"
	fmt.Printf("a = '%s'\n", a)
}
