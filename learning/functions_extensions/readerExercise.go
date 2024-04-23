package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// Read implements the io.Reader interface.
func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func mainInfAReader() {
	reader.Validate(MyReader{})
}
