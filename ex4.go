package main

import "fmt"

/**
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, …)
**/

//fbonacci calculate fibonacci
func fibonacci() func() int {
	total, nextTotal := 0, 1
	return func() int {
		result := total
		total, nextTotal = nextTotal, nextTotal+result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
