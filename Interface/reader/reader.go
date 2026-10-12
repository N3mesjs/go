package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var r io.Reader = strings.NewReader("Hello World!")
	buff := make([]byte, 2)

	for {
		n, err := r.Read(buff)

		if(n>0){
			fmt.Printf("Read %d byte: %q\n", n, buff[:n])
		}

		if(err == io.EOF){
			fmt.Println("End of the data stream")
			break
		}

		if(err != nil){
			fmt.Printf("An error occured %s\n", err)
			break
		}
	}
}
