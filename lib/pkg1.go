package golang101

import "fmt"

func private() {
	fmt.Println("Golang101 private function")
}

func Public() {
	fmt.Println("Golang101 public function")
	private()
}
