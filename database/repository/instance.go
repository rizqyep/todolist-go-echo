package repository

import "sync"

type repositoryPools struct {
	Todo TodoRepository
}

var repositoryInstance *repositoryPools
var once sync.Once

func InitRepositoryInstance() *repositoryPools {
	once.Do(func() {
		repositoryInstance = NewRepositoryInstance()
	})
	return repositoryInstance
}

func NewRepositoryInstance() *repositoryPools {
	return &repositoryPools{
		Todo: NewTodoRepository(),
	}
}
