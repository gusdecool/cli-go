package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "count lines")

	// Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// a scanner that will read input
	scanner := bufio.NewScanner(r)

	// define the scanner split to words
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	w := 0

	// count the word
	for scanner.Scan() {
		w++
	}

	// return the total
	return w
}
