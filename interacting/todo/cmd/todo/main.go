package main

import (
	"flag"
	"fmt"
	"github.com/gusdecool/cli-go/interacting/todo"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	// add CLI flag docs
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by a programmer\n", os.Args[0])
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}

	// parsing flag argument
	task := flag.String("task", "", "Task to be included in todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "item to be completed")
	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		// complete give item
		if err := l.Complete(*complete); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// add task
		l.Add(*task)

		// save the new list
		if err := l.Save(todoFileName); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// invalid flag provided
		_, _ = fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
