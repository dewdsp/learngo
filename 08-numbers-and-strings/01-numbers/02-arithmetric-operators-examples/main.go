package main

import "fmt"

func main() {
	fmt.Println("sum:", 3+2)
	fmt.Println("sum:", 2+3.5)
	fmt.Println("dif:", 3-1)
	fmt.Println("dif:", 3-0.5)
	fmt.Println("prod:", 4*5)
	fmt.Println("prod:", 5*2.5)
	fmt.Println("quot:", 8/2)
	fmt.Println("quot:", 8/1.5)

	fmt.Println("rem:", 8%3)

	f := 8.0
	fmt.Println("rem:", int(f)%3)
}
