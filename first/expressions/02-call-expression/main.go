package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPU: ", runtime.NumCPU())
}
