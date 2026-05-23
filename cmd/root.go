package cmd

import (
	"fmt"
	"go-task-tracker/task"
	"os"
)

func Execute(svc *task.Service) {
	if len(os.Args) < 2 {
		GeneralHelp()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		Add(svc, args)
	case "list":
		List(svc, args)
	case "help", "-h", "--help":
		Help(args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run 'task help' for usage.")
		os.Exit(1)
	}
}
