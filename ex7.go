package main

import "golang.org/x/tour/reader"

type MyReader struct{}

/**
Implement a Reader type that emits an infinite stream of the ASCII character ‘A’.
**/

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (i int, e error) {
	for x := range b {
		b[x] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
