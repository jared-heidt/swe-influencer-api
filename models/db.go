package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c connection) toString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("Error loading .env file: ", err.Error())
		return
	}

	connectionInfo := connection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err = sql.Open("postgres", connectionInfo.toString())

	if err != nil {
		fmt.Println("Error connecting to database", err.Error())
		return
	}

	fmt.Println("Database is open")

	err = db.Ping()

	if err != nil {
		fmt.Println("Error pinging database", err.Error())
	}

	fmt.Println("Database pinged successfully")
}
