package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var radius, area float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)
	area = 4 * math.Pi * radius * radius

	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}
