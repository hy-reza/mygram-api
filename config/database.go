package config

import (
	"fmt"

	"github.com/hy-reza/mygram-api/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {

	var dsn string

	if config.Env == "production" {
		dsn = config.DBUrl
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("Database connected successfuly!")
	helper.ErrorPanic(err)

	return db
}
