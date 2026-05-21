# Go Task Tracker
- A CLI task tracker built with zero external dependencies — pure Go standard library only.
- Tasks are persisted to a local JSON file.
- The architecture follows clean separation of concerns: CLI parsing → command handlers → domain logic → storage.

## Requirements
- [ ] Help page: Usage, Available commands, Flags
- [ ] Add, Update, Delete tasks
- [ ] Mark a task as in progress or done
- [ ] List all tasks
- [ ] List all tasks that are done
- [ ] List all tasks that are not done
- [ ] List all tasks that are in progress
- [ ] Generate autocompletion script for the specific shell

## Usage
### Commands
```sh
$ task help
Go Task Tracker — A simple CLI task manager

Usage:
  task <command> [arguments]

Commands:
  add <description>       Add a new task
  update <id> <text>      Update a task description
  delete <id>             Delete a task by ID
  mark <id> <status>      Change task status (todo, in-progress, done)
  list [status]           List all tasks, or filter by status
  help                    Show this help message

Run 'task help [command]' to learn more about a command
```

```sh
$ task add "buy groceries"
Task added: buy groceries
```

```sh
$ task list
ID  DESC            STATUS      UPDATED
--  ----            ------      -------
3   cook dinner     todo        an hour ago
2   organize table  done        last week
1   buy groceries   in-progress today
```

```sh
$ task list done
ID  DESC            STATUS      UPDATED
--  ----            ------      -------
2   organize table  done        last week
```

```sh
$ task list in-progress
ID  DESC            STATUS      UPDATED
--  ----            ------      -------
1   buy groceries   in-progress today
```

```sh
$ task list todo
ID  DESC            STATUS      UPDATED
--  ----            ------      -------
3   cook dinner     todo        an hour ago
```