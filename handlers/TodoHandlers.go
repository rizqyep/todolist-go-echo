package handlers

import (
	"echo-learn/entities"
	"echo-learn/models"
	"echo-learn/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	GetTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) TodoHandler {
	return &todoHandler{
		todoService: todoService,
	}
}

func parseId(c echo.Context) int64 {
	idParam := c.Param("id")
	idParse, err := strconv.Atoi(idParam)

	id := int64(idParse)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			ResponseStructure{
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "Something went wrong todo.",
			})
	}
	return id
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

func (handler *todoHandler) GetTodoById(c echo.Context) error {
	id := parseId(c)

	todo, err := handler.todoService.GetTodoById(id)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			ResponseStructure{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		ResponseStructure{
			Code:    http.StatusOK,
			Data:    todo,
			Message: "Sucessfully Fetch Todo!",
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
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "Something went wront while creating todo.",
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

func (handler *todoHandler) UpdateTodo(c echo.Context) error {
	var todo models.Todo
	id := parseId(c)

	err := (&echo.DefaultBinder{}).BindBody(c, &todo)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			ResponseStructure{
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "Something went wrong while updating todo.",
			})
	}

	updatedTodo := handler.todoService.UpdateTodo(id, todo)

	return c.JSON(
		http.StatusCreated,
		ResponseStructure{
			Code:    http.StatusCreated,
			Data:    updatedTodo,
			Message: "Successfully Updated Todos!",
		},
	)
}

func (handler *todoHandler) DeleteTodo(c echo.Context) error {
	id := parseId(c)

	handler.todoService.DeleteTodo(id)
	return c.JSON(
		http.StatusOK,
		ResponseStructure{
			Code:    http.StatusOK,
			Data:    nil,
			Message: fmt.Sprintf("Sucessfully deleted todo with ID %d", id),
		},
	)
}
