package main

import "fmt"

func double(arr []int) []int {
	for index, value := range arr {
		arr[index] = value * 2
	}
	return arr
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a = ", a)
	fmt.Println("a[0:] = ", a[0:])
	fmt.Println("a[0:len(a)] = ", a[0:len(a)])
	fmt.Println("a[5:] = ", a[5:])
	fmt.Println("a[0:len(a)] = ", a[0:len(a)])
	b := a[3:5]
	fmt.Println("b = a[3:5] = ", b)
	fmt.Println("len(b): ", len(b), "\tcap(b)", cap(b))
	fmt.Println("b[0] = ", b[0])
	fmt.Println("b[0:cap(b)] = ", b[0:cap(b)])

	fmt.Println("\nLets double a slice of a, a[3:5]")
	fmt.Println("double(a[3:5]) = ", double(a[3:5]))

	c := a[0:4]
	fmt.Println("C = a[0:4] = ", c)
	fmt.Println("***Wait a second, why is it 0 1 2 6 and not 0 1 2 3??")
	fmt.Println("***Arrays are copied when passed to a function, so are slices however a slice is really a data structure")

	fmt.Println("\nFinally lets set a[3] to 100")
	a[3] = 100
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)
	fmt.Println("C: ", c)

	fmt.Println("\nWe saw the copy() built-in, there is also append()")
	slice := []int{1, 2, 3}
	slice2 := []int{55, 66, 77}
	fmt.Println("Start slice: ", slice)
	fmt.Println("Start slice2:", slice2)

	// Add an item to a slice.
	slice = append(slice, 4)
	fmt.Println("Add one item:", slice)

	// Add one slice to another.
	slice = append(slice, slice2...)
	fmt.Println("Add one slice:", slice)

}
