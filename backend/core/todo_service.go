package core

type TodolistService interface {
	CreateTodo(Todo Todolist) error
	GetTodo() ([]Todolist, error)
	GetID(id int) (Todolist, error)
	Update(id int, Todo Todolist) (Todolist,error)
	DeleteID(id int) (Todolist,error)
}

type TodoServiceImpl struct {
	repo TodolistRepository
}

func NewTodolistService(repo TodolistRepository) TodolistService {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) CreateTodo(Todo Todolist) error {
	return s.repo.Save(Todo)
}

func (s *TodoServiceImpl) GetTodo() ([]Todolist, error) {
    return s.repo.Get()
}

func (s *TodoServiceImpl) GetID(id int) (Todolist,error){
	return s.repo.GetID(id)
}

func (s *TodoServiceImpl) DeleteID(id int) (Todolist,error) {
	return s.repo.DeleteID(id)
}

func (s *TodoServiceImpl) Update(id int,Todo Todolist) (Todolist,error) {
	return s.repo.Update(id,Todo)
}
