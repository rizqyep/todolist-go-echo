package handlers

import (
	"echo-learn/services"
	"sync"
)

type handlersPool struct {
	Todo TodoHandler
}

var handlersInstance *handlersPool
var once sync.Once
var serviceInstance = services.InitServiceInstance()

func InitHandlerInstance() *handlersPool {
	once.Do(func() {
		handlersInstance = NewHandlersInstance()
	})

	return handlersInstance
}

func NewHandlersInstance() *handlersPool {
	return &handlersPool{
		Todo: NewTodoHandler(serviceInstance.Todo),
	}
}
