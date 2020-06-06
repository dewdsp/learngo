package main

import "fmt"

func main() {
	fmt.Printf("%.2f\n", 8*3.14)
	fmt.Println(-(-2))
	fmt.Println(5 % 2)
	fmt.Println(2 - 3)

	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2
	fmt.Println(average)

	ratio := 1.0 / 10
	fmt.Printf("%.60f", ratio)
}
