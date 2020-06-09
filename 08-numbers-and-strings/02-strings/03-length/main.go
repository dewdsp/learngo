package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "สิทธการย์"

	fmt.Println(len(name))

	fmt.Println(utf8.RuneCountInString(name))
}
