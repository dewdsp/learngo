package main

import "fmt"

func main() {
	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? It is %v.\n", planet, hasLife)

	fmt.Printf(
		"%v is %v away, Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)

	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
}
