package main

import "fmt"

func main() {
	i, f, s, b := 314, 3.14, "Hello", true
	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)

	a, b := 14, true
	fmt.Println(a, b)

	a, c := 42, "good"
	fmt.Println(a, c)

	sum := 27 + 3.5
	fmt.Println(sum)

	d, _ := true, true
	fmt.Println(d)

	age, yourage := 35, 20
	ratio, age := 3.14, 42
	fmt.Println(age, yourage, ratio)
}
