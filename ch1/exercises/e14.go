// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
//Exercise 1.4 - Modify dup2 to print the names of all fiels in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]*DupLine)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.quantity > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.quantity, line, n.files)
		}
	}
}

func countLines(f *os.File, counts map[string]*DupLine) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, ok := counts[key]
		if ok {
			counts[input.Text()].quantity++
			counts[input.Text()].files += f.Name() + " "
		} else {
			counts[key] = new(DupLine)
			counts[input.Text()].quantity++
			counts[input.Text()].files += f.Name() + " "
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

type DupLine struct {
	quantity int
	files    string
}

//!-
