package main

import (
	"fmt"
	"math"
)

/**
Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
**/

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	tolerance := 1e-6

	for {
		oldz := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-oldz) < tolerance {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
