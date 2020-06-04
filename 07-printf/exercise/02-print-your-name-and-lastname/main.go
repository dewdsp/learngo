package main

import "fmt"

func main() {
	name, lastname := "Sitthakan", "Prasomsup"
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, name, lastname)
}
