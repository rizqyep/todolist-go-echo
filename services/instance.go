package services

import (
	"echo-learn/database/repository"
	"sync"
)

type servicesPool struct {
	Todo TodoService
}

var serviceInstance *servicesPool
var once sync.Once
var repositoryInstance = repository.InitRepositoryInstance()

func InitServiceInstance() *servicesPool {
	once.Do(func() {
		serviceInstance = NewServiceInstance()
	})
	return serviceInstance
}

func NewServiceInstance() *servicesPool {
	return &servicesPool{
		Todo: NewTodoService(repositoryInstance.Todo),
	}
}
