package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	strFields := strings.Fields(s);

	for _, key := range strFields{
		m[key]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
