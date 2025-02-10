package models

import "time"



type Task struct {
    ID          uint32 `db:"id"`
    Description string `db:"description"`
    Completed   bool   `db:"completed"`
    Priority    string `db:"priority"`
    AddedAt     time.Time `db:"created_at"`
    UpdatedAt   time.Time `db:"updated_at"`
}
