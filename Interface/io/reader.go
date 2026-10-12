package main

import (
	"fmt"
	"io"
	"strings"
)

/*
type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

func main() {
	var str string = "Hello, World!"
	b := make([]byte, 8)

	/*
	Returns a *strings.Reader, it is a constructor
	that implements the io.Reader interface!
	The read, read the bytes of the data and puts them 
	into the buffer, it returns the number of bytes read
	and when it finisces the file or string it
	return an error, EOF error(end of file)
	*/
	r := strings.NewReader(str)

	for {
		n, err := r.Read(b)
		fmt.Printf("Number of bytes read: %v, buffer: %v\n", n, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
