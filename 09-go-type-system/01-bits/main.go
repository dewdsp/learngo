package main

import "fmt"

func main() {
	fmt.Printf("%02b\n", 4)

	var b byte
	b = 2
	fmt.Printf("%08b = %d\n", b, b)

	b = 255
	fmt.Printf("%08b = %d\n", b, b)

}
