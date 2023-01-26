package database

import (
	"echo-learn/utils"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB
var once sync.Once

func NewConnection() *gorm.DB {

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	envMap := utils.GetEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", envMap["DB_HOST"], envMap["DB_USER"], envMap["DB_PASSWORD"], envMap["DB_DATABASE"], envMap["DB_PORT"])
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	return db
}

func GetConnection() *gorm.DB {
	once.Do(func() {
		dbInstance = NewConnection()
	})

	return dbInstance
}
