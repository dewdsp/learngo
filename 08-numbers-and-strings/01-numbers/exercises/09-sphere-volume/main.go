package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var volume, radius float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)
	volume = (4. / 3) * math.Pi * math.Pow(radius, 3.)

	fmt.Printf("radius: %g -> volume %.2f\n", radius, volume)

}
