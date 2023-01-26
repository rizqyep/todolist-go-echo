package services

import (
	"echo-learn/database/repository"
	"echo-learn/entities"
	"echo-learn/models"
)

type TodoService interface {
	GetTodos() *[]models.Todo
	GetTodoById(id int64) (*models.Todo, error)
	CreateTodo(todo entities.CreateTodoRequest) *models.Todo
	UpdateTodo(id int64, todo models.Todo) *models.Todo
	DeleteTodo(id int64)
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

func (service *todoService) GetTodoById(id int64) (*models.Todo, error) {
	return service.todoRepository.GetTodoById(id)
}

func (service *todoService) CreateTodo(todo entities.CreateTodoRequest) *models.Todo {
	newTodo := models.Todo{
		Task: todo.Task,
	}

	return service.todoRepository.CreateTodo(&newTodo)
}

func (service *todoService) UpdateTodo(id int64, todo models.Todo) *models.Todo {
	return service.todoRepository.UpdateTodo(id, todo)
}

func (service *todoService) DeleteTodo(id int64) {
	service.todoRepository.DeleteTodo(id)
}
