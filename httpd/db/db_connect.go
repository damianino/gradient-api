package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func PostgresConnect() error{
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"))
	var err error
	DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if DB.Ping() != nil {
		return fmt.Errorf("couldnt reach db")
	}
	fmt.Println("Connected to db!")
	return nil
}