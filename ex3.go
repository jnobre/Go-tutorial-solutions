package main

import (
	"golang.org/x/tour/wc"

	"strings"
)

/**
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.
Test function runs a test suite against the provided function and prints success or failure.
**/

//WordCount should return a map of the counts of each “word” in the string s
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
