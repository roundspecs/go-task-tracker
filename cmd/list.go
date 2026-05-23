package cmd

import (
	"fmt"
	"go-task-tracker/display"
	"go-task-tracker/task"
	"os"
)

func List(svc *task.Service, args []string) {
	if len(args) > 1 {
		fmt.Println("Error: list command takes only one filter")
		fmt.Println("Usage: task list")
		fmt.Println("   or, task list [status]")
		os.Exit(1)
	}

	var filter *task.Status

	if len(args) == 1 {
		s := task.Status(args[0])
		if s.IsValid() {
			filter = &s
		} else {
			fmt.Println("Error: invalid status")
			fmt.Println("Valid status: todo, in-progress, done")
			os.Exit(1)
		}
	}

	t, err := svc.List(filter)
	if err != nil {
		fmt.Printf("Error: listing tasks: %v\n", err)
		os.Exit(1)
	}

	display.Table(t)
}
