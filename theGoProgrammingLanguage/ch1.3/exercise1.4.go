// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counts map[string]int
	var files map[string]string
	files = make(map[string]string)
	counts = make(map[string]int)
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dpf2: %v\n", err)
			continue
		}
		countLines(f, counts, files)
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, files[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, files map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		files[input.Text()] = f.Name()
	}
}
