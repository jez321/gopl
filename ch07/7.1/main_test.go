package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "a b c", want: 3},
		{input: "ab c", want: 2},
		{input: "c", want: 1},
		{input: "", want: 0},
	}

	for _, tc := range tests {
		var wc WordCounter
		fmt.Fprint(&wc, tc.input)
		if wc != WordCounter(tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, wc)
		}
	}
}
