package main

import (
	"fmt"
	"go-task-tracker/cmd"
	"go-task-tracker/storage"
	"go-task-tracker/task"
	"os"
)

func main() {
	store, err := storage.NewJSONStorage(".task.json")
	if err != nil {
		fmt.Printf("Error: initializing storage %v", err)
		os.Exit(1)
	}

	svc := task.NewService(store)

	cmd.Execute(svc)
}
