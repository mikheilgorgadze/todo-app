package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikheilgorgadze/todo-app/models"
)

var DB *sqlx.DB


func InitDB(dbPath string) error {
    var err error
    DB, err = sqlx.Open("sqlite3", dbPath)
    if err!=nil {
        return err
    }

    err = DB.Ping()
    if err!=nil {
        return err
    }

    return nil
}

func CloseDB() {
    if DB!=nil {
        DB.Close()
    }
}

func AddTask(task *models.Task) error {
    query := `INSERT INTO tasks(
        completed,
        description,
        priority,
        updated_at
    ) values (?, ?, ?, ?)`
    _, err := DB.Exec(query, task.Completed, task.Description, task.Priority, task.UpdatedAt)
    if err!=nil {
        return err
    }

    return nil
}

func GetTasks() (*[]models.Task, error) {
    var t []models.Task

    err := DB.Select(&t, "select id, completed, description, priority, created_at, updated_at  from tasks t")
    if err!=nil {
        return nil, err
    }
    return &t, nil 
}

func MarkTaskCompleted(taskId int, completed bool) error{
    query := `UPDATE tasks
                 SET completed = ?
               WHERE id = ?`
    _, err := DB.Exec(query, completed, taskId)
    if err!=nil{
        return err
    }
    return nil
}

func ChangeTaskPriority(taskId int, priority string) error {
    query := `UPDATE tasks
                 SET priority = ?
               WHERE id = ?`
    _, err := DB.Exec(query, priority, taskId)
    if err!=nil{
        return err
    }
    return nil
}

func DeleteTask(taskId int) error {
    query := `DELETE FROM tasks WHERE id = ?`
    _, err := DB.Exec(query, taskId)
    if err!=nil{
        return err
    }
    return nil
}
