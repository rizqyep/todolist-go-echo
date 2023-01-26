package repository

import (
	"echo-learn/database"
	"echo-learn/models"
	"echo-learn/utils"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetTodos() *[]models.Todo
	GetTodoById(id int64) (*models.Todo, error)
	CreateTodo(todo *models.Todo) *models.Todo
	UpdateTodo(id int64, todo models.Todo) *models.Todo
	DeleteTodo(id int64)
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

func (repository *todoRepository) GetTodoById(id int64) (*models.Todo, error) {
	var todo models.Todo
	result := repository.db.First(&todo, id)
	if result.Error != nil {
		return nil, utils.HandleDBError(result.Error)
	}

	return &todo, nil

}

func (repository *todoRepository) CreateTodo(todo *models.Todo) *models.Todo {

	repository.db.Create(todo)

	return todo
}

func (repository *todoRepository) UpdateTodo(id int64, payload models.Todo) *models.Todo {

	var todo models.Todo
	repository.db.First(&todo, id)
	repository.db.Model(&todo).Updates(payload)
	return &todo
}

func (repository *todoRepository) DeleteTodo(id int64) {
	repository.db.Delete(&models.Todo{}, id)
}
