package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := "Test the product.   "
	name := strings.TrimRight(msg, " ")
	fmt.Println(utf8.RuneCountInString(name))
}
