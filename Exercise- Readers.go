package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (a MyReader) Read(b []byte) (int, error){
	b[0] = 'A'
	b[1] = 'A'
	return 2, nil
}

func main() {
	reader.Validate(MyReader{})
}
