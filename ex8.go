package main

import (
	"fmt"
	"strconv"
)

type foo struct {
	id int
}

type bar struct {
	id int
}

type foobar interface {
	Dump() string
}

func (f foo) Dump() string {
	return strconv.Itoa(f.id)
}

func (b bar) Dump() string {
	return strconv.Itoa(b.id)
}

func dumper(d foobar) {
	fmt.Println("Dump: ", d.Dump())
}

func main() {
	a := foo{1}
	b := bar{2}
	dumper(a)
	dumper(b)
}
