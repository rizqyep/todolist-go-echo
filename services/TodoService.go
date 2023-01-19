package services

import (
	"echo-learn/database/repository"
	"echo-learn/entities"
	"echo-learn/models"
)

type TodoService interface {
	GetTodos() *[]models.Todo
	CreateTodo(todo entities.CreateTodoRequest) *models.Todo
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (service *todoService) GetTodos() *[]models.Todo {
	return service.todoRepository.GetTodos()
}

func (service *todoService) CreateTodo(todo entities.CreateTodoRequest) *models.Todo {
	newTodo := models.Todo{
		Task: todo.Task,
	}

	return service.todoRepository.CreateTodo(&newTodo)

}
