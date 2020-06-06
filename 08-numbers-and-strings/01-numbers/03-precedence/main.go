package main

import "fmt"

func main() {
	n, m := 1, 5
	fmt.Println(2 + 1*m/n)
	fmt.Println(((2 + 1) * m) / n)

	fmt.Println(1 + 5 - 3*10/2)

	fmt.Println(
		(1 + 5 - 3) * (10 / 2),
	)

	celsius := 35.
	// correct formular = (9*celsius +160) / 5
	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%g C is %g F\n", celsius, fahrenheit)
}
