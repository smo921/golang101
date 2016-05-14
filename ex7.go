package main

import "fmt"
import "reflect"

type foo int

type bar struct {
	a    int
	b    float32
	c, d bool
	e    complex64
}

func main() {
	b := true
	fmt.Printf("b is a %T\n", b)
	fmt.Println("The reflect package says b is a ", reflect.TypeOf(b))

	var c foo
	fmt.Printf("\nc is a %T\n", c)
	fmt.Println("The reflect package says c is a ", reflect.TypeOf(c))

	var d int32
	e := 10
	f := 0.0
	fmt.Printf("\nd is a %T\n", d)
	fmt.Printf("e is a %T\n", e)
	fmt.Printf("f is a %T\n", f)

	mybar := bar{}
	fmt.Printf("\nmybar: %v\n", mybar)
	mybar.a = 1
	mybar.b = 2.0
	mybar.c = true
	mybar.e = complex(10.0, 5.0)
	fmt.Printf("mybar: %v\n", mybar)
}
