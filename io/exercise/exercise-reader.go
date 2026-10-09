package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(p []byte) (n int, err error){
	for i := range p {
		p[i] = 'A'
	}

	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
