# todolist-go-echo

Simple todolist rest api with Golang, implemented with 3 Layers Architecure.

HTTP Framework :<a href="https://echo.labstack.com/">Echo</a>
ORM : <a href="https://gorm.io/docs/">Gorm</a>
Because Gorm's migration doesnt have detailed migration log, i used `migrate` can check the ( <a href="github.com/golang-migrate/migrate">details here</a>


Also added singletons with once for resource optimization on each layer's instance 
(check `repository/instance.go` or `handlers/instance.go` or `service/instance.go`)
