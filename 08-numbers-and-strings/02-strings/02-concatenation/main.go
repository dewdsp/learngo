package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(
		"hello" + "," + "how" + " " + "are you" + " ",
	)

	fmt.Println(
		`hello` + "," + "how" + " " + "are you" + " ",
	)

	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)
}
