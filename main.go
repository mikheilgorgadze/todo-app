package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mikheilgorgadze/todo-app/database"
	"github.com/mikheilgorgadze/todo-app/task"
	"github.com/mikheilgorgadze/todo-app/ui"
)

func main() {

    t := task.NewTask()

    err := database.InitDB("todo.db")
    if err!=nil {
        fmt.Println("Error connecting database: ", err.Error())
    }
    defer database.CloseDB()

    err = database.RunMigrations()
    if err!=nil{
        fmt.Println(err.Error())
    }

	for {
		ui.ShowMenu()
		option := getUserInput("Enter your choice: ")

		switch option {
		case "1":
			ui.ShowTasks(t.GetTasks())
		case "2":
            t.AddTask()
		case "3":
			ui.ShowTasks(t.GetTasks())
            t.MarkTaskCompleted()
		case "4":
            ui.ShowTasks(t.GetTasks())
            t.ChangeTaskPriority()
        case "5":
            ui.ShowTasks(t.GetTasks())
            t.DeleteTask()
		case "6":
			fmt.Println("Exiting the todo application.")
		    return
        default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}



func getUserInput(s string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
