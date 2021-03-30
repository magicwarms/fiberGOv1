package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/magicwarms/fiberGOv1/models"
)

var (
	// DB is a exported connection
	DB *gorm.DB
)

// InitDatabase is initial Setup for DB Connection
func InitDatabase() {
	var err error
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

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	DB.AutoMigrate(models.Books{})
	DB.AutoMigrate(models.Authors{})
	fmt.Println("Connection Opened to Database")
}
