package ui

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/mikheilgorgadze/todo-app/models"
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
	fmt.Println("5. Delete Task")
	fmt.Println("6. Exit program")
}


func ShowTasks(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	table := uitable.New()
	fmt.Println("Tasks")
	fmt.Println("--------------------------------")
	table.AddRow("Task No.", "Status", "Priority Level", "Description", "Task Added At", "Task Updated At")
	for _, task := range tasks {
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
		table.AddRow(task.ID, status, task.Priority, task.Description, task.AddedAt, task.UpdatedAt)
	}
    fmt.Println(table)
}

