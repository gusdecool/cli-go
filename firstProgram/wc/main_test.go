package main

import (
	"bytes"
	"testing"
)

// test count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")

	exp := 4
	res := count(b, Words)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

// test check lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res := count(b, Lines)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}

// test check bytes
func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")

	exp := 23
	res := count(b, Bytes)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}
