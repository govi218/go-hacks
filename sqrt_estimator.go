package main

import (
	"fmt"
	// "math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 10; i++ {
		if z*z == x {
			return z
		}
		fmt.Println("Guess #", i + 1, ": ", z)
		z -= (z*z - x) / (2*z)
	}
	return
}

func main() {
	fmt.Println(Sqrt(999))
}