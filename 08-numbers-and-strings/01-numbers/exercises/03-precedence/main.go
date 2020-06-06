package main

import "fmt"

func main() {
	fmt.Println(10 + 5 - (5 - 10))     // result should be 20
	fmt.Println(-10 + 0.5 - (1 + 5.5)) // result should be -16
	fmt.Println(5 + 10*(2-5))          // result should be -25
	fmt.Println(0.5 * (2 - 1))         // result should be 0.5
	fmt.Println((3+1)/2*10 + 4)        // result should be 24
	fmt.Println(10 / 2 * (10 % 7))     // result should be 15
	fmt.Println(100 / (5. / 2))        // result should be 40
}
