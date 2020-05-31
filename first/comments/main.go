package main

import (
	"fmt"
	"runtime"
)

/**
main function
Go execution this program
*/
func main() {
	fmt.Println("Hello!!!")

	fmt.Println("Number of CPU:", runtime.NumCPU())
}
