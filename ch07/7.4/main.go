package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type stringReader string

func (r *stringReader) Read(p []byte) (int, error) {
	copy(p, []byte(*r))
	return len(*r), io.EOF
}

func NewReader(s string) io.Reader {
	sr := stringReader(s)
	return &sr
}

func main() {
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}
