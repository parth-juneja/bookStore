package config

import (
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
}

func DataBaseConnection() *gorm.DB {
	fmt.Println("connecting to database...")
	port, _ := strconv.Atoi(GetEnv("DB_PORT", "5432"))
	config := DBConfig{
		Host:     GetEnv("DB_HOST", "localhost"),
		User:     GetEnv("DB_USER", "user"),
		Password: GetEnv("DB_PASSWORD", "password"),
		DBName:   GetEnv("DB_NAME", "books"),
		Port:     port,
		SSLMode:  "disable",
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	//FIXME: error here
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
