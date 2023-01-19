package routes

import (
	"echo-learn/handlers"

	"github.com/labstack/echo/v4"
)

func RouteHandlers(r *echo.Echo) {
	handlers := handlers.InitHandlerInstance()
	r.GET("/todos", handlers.Todo.GetTodos)
	r.POST("/todos", handlers.Todo.CreateTodo)

}
