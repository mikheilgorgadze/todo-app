package database

import "fmt"


func RunMigrations() error {
    _, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            completed INTEGER NOT NULL,
            description TEXT NOT NULL,
            priority TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        )
        `)
    if err!=nil {
        return fmt.Errorf("migration failed: %w", err)
    }

    return nil
}
