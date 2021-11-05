package main

import (
	"fmt"
	"io"
)

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	written, err := b.w.Write(p)
	if err != nil {
		return len(p), err
	}
	b.count += int64(written)
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{w, 0}
	return &bc, &bc.count
}

func main() {
	writer, count := CountingWriter(io.Discard)
	fmt.Fprint(writer, "HELLO")
	fmt.Println(*count)
	fmt.Fprint(writer, "ABCDEFGHIJK")
	fmt.Println(*count)
}
