package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		Add(args)
	}
}