package ui

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/mikheilgorgadze/todo-app/task"
)


func ColorCodePriority(priority *string) {
	switch *priority {
    case "Low":
        *priority = color.YellowString("Low")
    case "Medium":
        *priority = color.GreenString("Medium")
    case "High":
        *priority = color.CyanString("High") 
	case "undefined":
		*priority = color.RedString("undefined")
	}
}

func ShowMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add Tasks")
	fmt.Println("3. Mark Task as Completed")
	fmt.Println("4. Change Task Priority")
	fmt.Println("5. Save Tasks to File")
	fmt.Println("6. Exit program")
}


func ShowTasks(tasks []task.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	table := uitable.New()
	fmt.Println("Tasks")
	fmt.Println("--------------------------------")
	table.AddRow("Task No.", "Status", "Priority Level", "Text")
	for i, task := range tasks {
		status := "Not Completed"
		if task.Completed {
            status = "Completed"
        }
		if task.Priority == "" {
			task.Priority = "undefined"
			ColorCodePriority(&task.Priority)
		} else {
			ColorCodePriority(&task.Priority)
		}
		table.AddRow(i+1, status, task.Priority, task.Text)
	}
    fmt.Println(table)
}

