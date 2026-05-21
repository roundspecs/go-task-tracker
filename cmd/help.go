package cmd

import "fmt"

// Help prints the usage instructions and available commands.
func Help() {
	fmt.Println("Go Task Tracker — A simple CLI task manager")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add <description>       Add a new task")
	fmt.Println("  update <id> <text>      Update a task description")
	fmt.Println("  delete <id>             Delete a task by ID")
	fmt.Println("  mark <id> [status]      Change task status (todo, in-progress, done)")
	fmt.Println("  list [status]           List all tasks, or filter by status")
	fmt.Println("  help                    Show this help message")
	fmt.Println()
	fmt.Println("Run 'task help [command]' to learn more about a command")
}