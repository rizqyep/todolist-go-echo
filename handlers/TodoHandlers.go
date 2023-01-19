package handlers

import (
	"echo-learn/entities"
	"echo-learn/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	GetTodos(c echo.Context) error
	CreateTodo(c echo.Context) error
}

type todoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) TodoHandler {
	return &todoHandler{
		todoService: todoService,
	}
}

func (handler *todoHandler) GetTodos(c echo.Context) error {
	todos := handler.todoService.GetTodos()
	return c.JSON(
		http.StatusOK,
		ResponseStructure{
			Code:    http.StatusOK,
			Data:    todos,
			Message: "Successfully Fetch Todos!",
		},
	)
}

func (handler *todoHandler) CreateTodo(c echo.Context) error {
	var todo entities.CreateTodoRequest
	err := c.Bind(&todo)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			ResponseStructure{
				Code:    http.StatusOK,
				Data:    nil,
				Message: "Successfully Fetch Todos!",
			},
		)
	}

	newTodo := handler.todoService.CreateTodo(todo)

	return c.JSON(
		http.StatusCreated,
		ResponseStructure{
			Code:    http.StatusCreated,
			Data:    newTodo,
			Message: "Successfully Create Todos!",
		},
	)
}
