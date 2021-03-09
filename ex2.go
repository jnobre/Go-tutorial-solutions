package main

import "golang.org/x/tour/pic"

/**
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
**/

//Pic It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
func Pic(dx, dy int) [][]uint8 {
	var result = make([][]uint8, dy)

	for x := range result {
		result[x] = make([]uint8, dx)
		for y := range result[x] {
			result[x][y] = uint8(x * y)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
