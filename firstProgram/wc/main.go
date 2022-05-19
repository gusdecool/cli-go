package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	Words int = iota
	Lines
	Bytes
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("b", false, "count bytes, used first")

	check := Words

	switch true {
	case *bytes:
		check = Bytes
		break
	case *lines:
		check = Lines
		break
	}

	// Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, check))
}

func count(r io.Reader, check int) int {
	// a scanner that will read input, by default lines
	scanner := bufio.NewScanner(r)

	// define the scanner split to words
	switch check {
	case Words:
		scanner.Split(bufio.ScanWords)
		break
	case Bytes:
		scanner.Split(bufio.ScanBytes)
		break
	case Lines:
		scanner.Split(bufio.ScanLines)
		break
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
