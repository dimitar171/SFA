package main

import (
	"io"
	"os"
)

type ReverseStringReader struct {
	read string
}

//constructor
func NewReverseStringReader(input string) *ReverseStringReader {
	return &ReverseStringReader{input}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//implement IO READER
func (r *ReverseStringReader) Read(p []byte) (int, error) {
	pom := []byte(r.read)
	myString := string(pom[:])
	result := Reverse(myString)
	n := copy(p, result)
	return n, io.EOF
}
func main() {

	ad := NewReverseStringReader("apple")
	io.Copy(os.Stdout, ad)
}
