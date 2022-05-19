package main

import (
	"fmt"
	"github.com/gusdecool/cli-go/interacting/todo"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var item string

	switch {
	case len(os.Args) == 1:
		// no argument, print all task
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		// concatenate all arguments into string
		item = strings.Join(os.Args[1:], " ")
	}

	l.Add(item)

	// save the new list
	if err := l.Save(todoFileName); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
