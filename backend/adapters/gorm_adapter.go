package adapters

import (
	"go/clean/core"
	"gorm.io/gorm"
)

type GormTodoRepository struct {
	db *gorm.DB
}



func NewGormOrderRepository(db *gorm.DB) core.TodolistRepository {
	return &GormTodoRepository{db: db}
}

func (r * GormTodoRepository) Save(Todo core.Todolist) error {
	if result := r.db.Save(&Todo) ; result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormTodoRepository) Get() ([]core.Todolist, error) {
    var todos []core.Todolist
    result := r.db.Find(&todos)
    return todos, result.Error
}
func (r *GormTodoRepository) GetID(id int) (core.Todolist,error){
	var todo core.Todolist
	result := r.db.First(&todo,id)
	return todo,result.Error
}
func (r *GormTodoRepository) DeleteID(id int) (core.Todolist,error) {
	var todo core.Todolist
	result := r.db.Delete(&todo,id)
	return todo,result.Error
}

func (r *GormTodoRepository) Update(id int, todo core.Todolist) (core.Todolist, error) {
    var updatedTodo core.Todolist
    result := r.db.Model(&core.Todolist{}).Where("id = ?", id).Updates(todo)
    return updatedTodo, result.Error
}