package main

import (
	"fmt"
	"time"
)

type gram float64
type ounce float64

func main() {
	type Duration int64

	var ms int64 = 1000
	var ns Duration

	ns = Duration(ms)
	ms = int64(ns)

	h, _ := time.ParseDuration("4h30m")

	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))

	var m int64 = 2

	h *= time.Duration(m)
	fmt.Println(h)

	var g gram = 1000
	var o ounce
	o = ounce(g) * 0.035274
	fmt.Printf("%g grams is %.2f ounces\n", g, o)

}
