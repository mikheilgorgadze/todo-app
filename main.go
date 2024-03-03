package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mikheilgorgadze/todo-app/task"
	"github.com/mikheilgorgadze/todo-app/ui"
)

func main() {
	tasks := []task.Task{}
	tasks, err := readTaskFromFile("tasks.txt")

	if err!=nil{
		fmt.Println("Error occured: ", err)
	}
	for {
		ui.ShowMenu()
		option := getUserInput("Enter your choice: ")

		switch option {
		case "1":
			ui.ShowTasks(tasks)
		case "2":
			addTask(&tasks)
		case "3":
			markTaskCompleted(&tasks)
		case "4":
			changeTaskPriority(&tasks)
		case "5":
			saveTasksToFile(tasks)
		case "6":
			fmt.Println("Exiting the todo application.")
		    return
        default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
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


func addTask(tasks *[]task.Task){
	taskText := getUserInput("Enter task description: ")
	*tasks = append(*tasks, task.Task{Text:taskText})
	fmt.Println("Task added successfully.")
}

func markTaskCompleted(tasks *[]task.Task){
	ui.ShowTasks(*tasks)
	taskIndexStr := getUserInput("Enter task number to mark as completed: ")
	taskIndex, err := strconv.Atoi(taskIndexStr)
	if err!=nil || taskIndex < 1 || taskIndex > len(*tasks) {
		fmt.Println("Invalid task index, please try again")
		return
	}
	(*tasks)[taskIndex-1].Completed = true
	fmt.Println("Task marked as completed")
}

func changeTaskPriority(tasks *[]task.Task) {
	ui.ShowTasks(*tasks)
	taskIndexStr := getUserInput("Enter task number to change its priority: ")
	taskIndex, err := strconv.Atoi(taskIndexStr)
	if err!=nil || taskIndex < 1 || taskIndex > len(*tasks) {
        fmt.Println("Invalid task index, please try again")
        return
    }

	showTaskPriorities()
	priorityStr := getUserInput("Enter new priority number: ")
	priorityKey, userInputErr := strconv.Atoi(priorityStr)

	if userInputErr!=nil  {
        fmt.Println("Invalid priority, please try again")
        return
    }

	priority, priorityKeyErr := decodePriorities(priorityKey)
	if priorityKeyErr!=nil {
		fmt.Println("Invalid priority number, please try again")
		return 
	}

	(*tasks)[taskIndex - 1].Priority = priority;
	fmt.Println("Task priority added.")
}

func showTaskPriorities(){
	fmt.Println("\nTask priorities")
	fmt.Println("1. Low (1)")
	fmt.Println("2. Medium (2)")
	fmt.Println("3. High (3)")
}

func saveTasksToFile(tasks []task.Task) {
	file, err := os.Create("tasks.txt")
	if err!= nil {
        fmt.Println("Error creating file: ", err)
        return
    }

	defer file.Close()
	for _, task := range tasks {
		if task.Priority == "" {
            task.Priority = "undefined"
        }
		file.WriteString(fmt.Sprintf("%v, Priority Level - %s, %s\n", task.Completed, task.Priority, task.Text))
	}
	fmt.Println("Task saved to file 'tasks.txt'")
}

func readTaskFromFile(fileName string) (tasks []task.Task, err error) {
	file, err := os.Open(fileName)
	if err!= nil {
        return tasks, err
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		task, err := parseTaskFromLine(line)
		if err!= nil {
			fmt.Println("error parsing task")
			continue
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func parseTaskFromLine(line string) (task task.Task, err error) {

	taskValues := strings.Split(line, ", ")

	if strings.HasPrefix(line, "true") {
		task.Completed = true
	} else {
		task.Completed = false;
	}


	for _, v := range taskValues {
		if strings.HasPrefix(v, "Priority Level - ") {
			task.Priority = strings.TrimPrefix(v, "Priority Level - ")
		} else {
			task.Text = strings.TrimSpace(v)
		}
	}

	return task, nil
}