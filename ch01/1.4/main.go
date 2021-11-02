// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type duplicate struct {
	count int
	files []string
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]duplicate)
		countLines(os.Stdin, counts, "")
		for line, n := range counts {
			if n.count > 1 {
				fmt.Printf("%d\t%s\n", n.count, line)
			}
		}
	} else {
		counts := make(map[string]duplicate)
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
		for line, n := range counts {
			if n.count > 1 {
				fmt.Printf("%d\t%s\t%s\n", n.count, line, strings.Join(n.files, " "))
			}
		}
	}
}

func countLines(f *os.File, counts map[string]duplicate, fn string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if text == "" {
			break
		}
		textEntry := counts[input.Text()]
		textEntry.count++
		if fn != "" {
			textEntry.files = append(textEntry.files, fn)
		}
		counts[input.Text()] = textEntry
	}
	// NOTE: ignoring potential errors from input.Err()
}
