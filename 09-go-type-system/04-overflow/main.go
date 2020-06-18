package main

import (
	"fmt"
	"math"
)

func main() {
	// int8(math.MaxInt8 + 1)
	var n int8 = math.MaxInt8
	fmt.Println("max int8    :", n)
	fmt.Println("max int8 +1 :", n+1)

	n = math.MinInt8
	fmt.Println("min int8    :", n)
	fmt.Println("min int8 -1 :", n-1)

	var un uint8
	fmt.Println("max uint8     :", un)
	fmt.Println("max uint8 - 1 :", un-1)

}
