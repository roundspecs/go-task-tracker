package display

import (
	"fmt"
	"go-task-tracker/task"
	"strconv"
	"strings"
	"time"
)

func Table(tasks []task.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fields := [...]string{"ID", "DESC", "STATUS", "CREATED"}
	fieldLen := [...]int{2, 4, 6, 6}

	for _, t := range tasks {
		fieldLen[0] = max(fieldLen[0], len(strconv.Itoa(t.ID)))
		fieldLen[1] = max(fieldLen[1], len(t.Description))
		fieldLen[2] = max(fieldLen[2], len(t.Status))
		fieldLen[3] = max(fieldLen[3], len(toReadableTime(t.CreatedAt)))
	}
	format := fmt.Sprintf("%%-%vv %%-%vv %%-%vv %%-%vv\n", fieldLen[0], fieldLen[1], fieldLen[2], fieldLen[3])

	fmt.Printf(format, fields[0], fields[1], fields[2], fields[3])
	for _, l := range fieldLen {
		fmt.Print(strings.Repeat("-", l) + " ")
	}
	fmt.Println()
	for _, t := range tasks {
		fmt.Printf(format, t.ID, t.Description, t.Status, toReadableTime(t.CreatedAt))
	}
}

func toReadableTime(t time.Time) string {
	delta := time.Since(t)
	if delta < time.Second*5 {
		return "just now"
	}
	if delta < time.Minute {
		return fmt.Sprintf("%.0fs ago", delta.Seconds())
	}
	if delta < time.Hour {
		return fmt.Sprintf("%.0fm ago", delta.Minutes())
	}
	if delta < time.Hour*24 {
		return fmt.Sprintf("%.0fh ago", delta.Hours())
	}
	if delta < time.Hour*24*7 {
		return fmt.Sprintf("last %v", strings.ToLower(t.Weekday().String()))
	}
	return fmt.Sprintf("%v, %v", strings.ToLower(t.Month().String()), t.Day())
}
