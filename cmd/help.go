package cmd

import "fmt"

// Help prints the usage instructions and available commands.
func Help(args []string) {
	if len(args) == 0 {
		GeneralHelp()
		return
	}

	command := args[0]
	switch command {
	case "add":
		HelpAdd()
	case "help":
		HelpHelp()
	default:
		fmt.Println("Unknown command: ", command)
		GeneralHelp()
	}
}

func GeneralHelp() {
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
	fmt.Println("  help [command]          Show help for a specific command")
	fmt.Println()
}

func HelpAdd() {
	fmt.Println("Add task")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task add <description>")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  task add \"buy groceries\"")
	fmt.Println("  task add buy groceries")
	fmt.Println()
}

func HelpHelp() {
	fmt.Println("Help")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task help")
	fmt.Println("  task help <command>")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  task help")
	fmt.Println("  task help add")
	fmt.Println()
}
