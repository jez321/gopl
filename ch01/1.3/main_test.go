// Echo1 prints its command-line arguments.
package main

import (
	"os"
	"strings"
	"testing"
)

// testStringsのlenが2以下の場合Loopが早い、3以上の場合StringJoinが早い

var testStrings = []string{"a", "b", "c", "d", "e", "f", "g"}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(os.Args, " ")
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for _, arg := range testStrings {
			s = s + sep + arg
			sep = " "
		}
		_ = s
	}
}
