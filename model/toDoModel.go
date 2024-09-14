package model

import "time"

type ToDoModel struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	IsCompleted bool      `db:"is_completed"`
}
