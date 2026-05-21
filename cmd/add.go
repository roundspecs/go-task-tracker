package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Add(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: task description is required")
		fmt.Println("Usage: task add \"buy groceries\"")
		fmt.Println("   or, task add buy groceries")
		os.Exit(1)
	}

	description := strings.Join(args, " ")

	fmt.Println("Description:", description)
}
