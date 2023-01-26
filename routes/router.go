package routes

import (
	"echo-learn/handlers"

	"github.com/labstack/echo/v4"
)

func RouteHandlers(r *echo.Echo) {
	handlers := handlers.InitHandlerInstance()
	r.GET("/todos", handlers.Todo.GetTodos)
	r.GET("/todos/:id", handlers.Todo.GetTodoById)
	r.POST("/todos", handlers.Todo.CreateTodo)
	r.PUT("/todos/:id", handlers.Todo.UpdateTodo)
	r.DELETE("/todos/:id", handlers.Todo.DeleteTodo)

}
