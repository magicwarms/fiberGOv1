package config

import (
	"fmt"
	"log"
	"os"
	"time"

	entity "fiberGOv1/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// initial db connection
var (
	DBConn *gorm.DB
)

// InitDatabase is initial Setup DB Connection
func InitDatabase() {
	dsn := "host=" + GoDotEnvVariable("DB_HOST") + " user=" + GoDotEnvVariable("DB_USERNAME") + " password=" + GoDotEnvVariable("DB_PASSWORD") + " dbname=" + GoDotEnvVariable("DB_DATABASE") + " port=" + GoDotEnvVariable("DB_PORT") + " sslmode=disable TimeZone=" + GoDotEnvVariable("TZ")
	appEnv := GoDotEnvVariable("APP_ENV")
	logLvl := logger.Info
	if appEnv == "production" {
		logLvl = logger.Silent
	}
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLvl,      // Log level
			Colorful:      true,        // Enable color
		},
	)
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	dbConnection.AutoMigrate(entity.Book{})
	fmt.Println("Connection Opened to Database")
}
