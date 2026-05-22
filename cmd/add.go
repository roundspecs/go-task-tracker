package cmd

import (
	"fmt"
	"go-task-tracker/task"
	"os"
	"strings"
)

func Add(svc *task.Service, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: task description is required")
		fmt.Println("Usage: task add \"buy groceries\"")
		fmt.Println("   or, task add buy groceries")
		os.Exit(1)
	}

	description := strings.Join(args, " ")

	t, err := svc.Add(description)
	if err != nil {
		fmt.Printf("Error: adding task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task#%v created: %v\n", t.ID, t.Description)
}
