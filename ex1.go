package main

import (
	"fmt"
	"math"
)

/**
Given a number x, we want to find the number z for which z² is most nearly x.
**/

//Sqrt Calculate the square root
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
