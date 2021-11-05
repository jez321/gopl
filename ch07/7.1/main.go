package main

import "strings"

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	words := strings.Split(string(p), " ")
	*w += WordCounter(len(words))
	return len(p), nil
}

func main() {

}
