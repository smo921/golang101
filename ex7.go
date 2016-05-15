package main

import "fmt"
import "reflect"

type bar struct {
	a    int
	b    float32
	c, d bool
	e    complex64
}

func (b bar) dump() {
	fmt.Printf("a: %v\nb: %v\nc: %v\nd: %v\ne: %v\n", b.a, b.b, b.c, b.d, b.e)
}

func main() {
	b := true
	fmt.Printf("b is a %T\n", b)
	fmt.Println("The reflect package says b is a ", reflect.TypeOf(b))

	type foo int
	var c foo
	fmt.Printf("\nc is a %T\n", c)
	fmt.Println("The reflect package says c is a ", reflect.TypeOf(c))

	var d int32
	e := 10
	f := 0.0
	fmt.Printf("\nd is a %T\n", d)
	fmt.Printf("e is a %T\n", e)
	fmt.Printf("f is a %T\n", f)

	var mymap map[string]string
	fmt.Println("\nMyMap: ", mymap)
	mymap["this"] = "will crash"

	mymap = make(map[string]string)
	mymap["foo"] = "bar"
	mymap["wom"] = "bat"
	fmt.Println("MyMap: ", mymap)

	mybar := bar{}
	fmt.Printf("\nmybar: %v\n", mybar)
	mybar.a = 1
	mybar.b = 2.0
	mybar.c = true
	mybar.e = complex(10.0, 5.0)
	fmt.Printf("mybar: %v\n", mybar)

	fmt.Println("\nmybar.dump:")
	mybar.dump()
}
