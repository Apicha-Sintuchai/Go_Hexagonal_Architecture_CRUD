package core

import "time"

type Todolist struct {
	ID          uint
	Work        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
