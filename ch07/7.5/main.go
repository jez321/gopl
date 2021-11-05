package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r     io.Reader
	limit int64
	read  int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if lr.read >= lr.limit {
		return 0, io.EOF
	}
	if int64(len(p))+lr.read < lr.limit {
		lr.read += int64(len(p))
		return lr.r.Read(p)
	}
	readLimit := lr.limit - lr.read
	toRead := p[:readLimit]
	n, err = lr.r.Read(toRead)
	lr.read += int64(n)
	if err != nil {
		return n, err
	}
	return n, io.EOF
}

func LimitReader(r io.Reader, n int64) io.Reader {
	reader := limitReader{r, n, 0}
	return &reader
}

func main() {
	sr := strings.NewReader("this is some text")
	lr := LimitReader(sr, 8)

	b := &bytes.Buffer{}
	n, err := b.ReadFrom(lr)

	fmt.Println(n, err)
}
