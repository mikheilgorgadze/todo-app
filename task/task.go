package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mikheilgorgadze/todo-app/database"
	"github.com/mikheilgorgadze/todo-app/models"
)

type TaskController struct {
}

func NewTask() *TaskController{
    return &TaskController{}
}

func (t *TaskController) GetTasks() []models.Task{
    tasks, err := database.GetTasks()
    if err!=nil {
        fmt.Println("Error reading tasks: ", err.Error()) 
        return nil
    }

    return *tasks
}


func (t *TaskController) AddTask() {
	taskText := getUserInput("Enter task description: ")
    priority, err := getPriorityFromUser()
    if err != nil {
        fmt.Println(err.Error())
    }
    task := &models.Task{
        Description: taskText,
        Completed: false,
        Priority: priority,
        UpdatedAt: time.Now(),
    }
    err = database.AddTask(task)
    if err!=nil {
        fmt.Println(err.Error())
    }


	fmt.Println("Task added successfully.")
}

func (t *TaskController) MarkTaskCompleted(){
    taskIndex, err := getTaskIndex("Enter task index to mark it completed: ")
    if err!=nil {
        return
    }

    err = database.MarkTaskCompleted(taskIndex, true)
    if err!=nil {
        fmt.Println(err.Error())
        return
    }

	fmt.Println("Task marked as completed")
}

func (t *TaskController) ChangeTaskPriority() {
    taskIndex, err := getTaskIndex("Enter task index to change it's priority: ")
    if err!=nil {
        return
    }

    priority, err := getPriorityFromUser()
    if err!=nil {
        fmt.Println(err.Error())
    }

    err = database.ChangeTaskPriority(taskIndex, priority)
    if err!=nil {
        fmt.Println(err.Error())
        return
    }

	fmt.Println("Task priority added.")
}

func (t *TaskController) DeleteTask() {
    taskIndex, err := getTaskIndex("Enter task index to delete: ")
    if err!=nil{
        return
    }

    err = database.DeleteTask(taskIndex)
    if err!=nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("Task deleted successfully.")
}

func getPriorityFromUser() (priority string, err error) {
    showTaskPriorities()
	priorityStr := getUserInput("Enter new priority number: ")
	priorityKey, userInputErr := strconv.Atoi(priorityStr)

	if userInputErr!=nil  {
        return "", fmt.Errorf("Invalid priority, please try again") 
    }

	priority, priorityKeyErr := decodePriorities(priorityKey)
	if priorityKeyErr!=nil {
        return "", fmt.Errorf("Invalid priority, please try again") 
	}
    
    return priority, nil
}

func showTaskPriorities(){
	fmt.Println("\nTask priorities")
	fmt.Println("1. Low (1)")
	fmt.Println("2. Medium (2)")
	fmt.Println("3. High (3)")
}

func decodePriorities(priorityKey int) (priority string, err error) {
	var priorities = map[int] string{
		1: "Low",
		2: "Medium",
		3: "High",
	}

	priority, ok := priorities[priorityKey]
	if !ok {
		err = fmt.Errorf("no priority found with this key")
	}
	return priority, err
}

func getUserInput(s string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getTaskIndex(prompt string) (int, error) {
	taskIndexStr := getUserInput(prompt)
	taskIndex, err := strconv.Atoi(taskIndexStr)
	if err!=nil || taskIndex < 1 {
        fmt.Println("Invalid task index, please try again")
        return 0, err
    }
    return taskIndex, nil
}
