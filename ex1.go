package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func hello() {
	fmt.Println("Hello World")
}

func randomError() (string, error) {
	r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	if r.Intn(2) == 0 {
		return "", fmt.Errorf("An error occured")
	}
	return "You got lucky, no error this time", nil
}

func goodbye() string {
	return "Goodbye"
}

func main() {
	defer fmt.Println(", cruel world")
	defer fmt.Print(goodbye())
	hello()
	foo, err := randomError()
	if err != nil {
		log.Println("Error: ", err)
	} else {
		fmt.Println(foo)
	}
}
