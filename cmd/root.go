package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		Help()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		Add(args)
	case "help", "-h", "--help":
		Help()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run 'task help' for usage.")
		os.Exit(1)
	}
}