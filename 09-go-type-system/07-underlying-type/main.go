package main

import (
	"fmt"

	weights "github.com/dewdsp/learngo/09-go-type-system/07-underlying-type/weight"
)

func main() {
	type (
		// Gram is underlying type is int
		Gram int
		// Kilogram is underlying type is int
		Kilogram Gram
		// Ton is underlying type is int
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Gram(int(apples)))
	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(weights.Gram(100))
	fmt.Printf("Type of weight.Gram: %T\n", weights.Gram(1))
	fmt.Printf("Type of weight.Gram: %T\n", Gram(1))

	type myGram weights.Gram
	var pepper myGram = 20
	fmt.Printf("Type of peper: %T\n", pepper)

}
