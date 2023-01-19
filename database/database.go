package database

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var once sync.Once

func NewConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=pgpass123 dbname=todos_go port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}

func GetConnection() *gorm.DB {
	once.Do(func() {
		dbInstance = NewConnection()
	})

	return dbInstance
}
