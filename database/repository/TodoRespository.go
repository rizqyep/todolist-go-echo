package repository

import (
	"echo-learn/database"
	"echo-learn/models"
	"log"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetTodos() *[]models.Todo
	CreateTodo(todo *models.Todo) *models.Todo
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{
		db: database.GetConnection(),
	}
}

func (repository *todoRepository) GetTodos() *[]models.Todo {
	var todos []models.Todo

	repository.db.Find(&todos)
	return &todos
}

func (repository *todoRepository) CreateTodo(todo *models.Todo) *models.Todo {

	result := repository.db.Create(todo)
	if result.Error != nil {
		log.Fatal("Error creating book")
	}
	newTodo := models.Todo{ID: todo.ID}
	repository.db.First(&newTodo)

	return &newTodo
}
