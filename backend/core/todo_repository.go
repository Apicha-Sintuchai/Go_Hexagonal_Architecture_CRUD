package core

type TodolistRepository interface {
	Save(Todo Todolist) error
	Get() ([]Todolist, error)
	GetID(id int) (Todolist,error)
	Update(id int, Todo Todolist) (Todolist,error)
	DeleteID(id int) (Todolist,error)
}
