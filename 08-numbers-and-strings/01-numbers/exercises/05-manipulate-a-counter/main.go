package main

import "fmt"

func main() {
	counter := 0

	counter++
	counter--
	counter += 5
	counter *= 10
	counter /= 5

	fmt.Println(counter)

}
