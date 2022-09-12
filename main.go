package main

import (
	"fmt"
	"gradientApi/httpd"
	"gradientApi/httpd/db"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	rand.Seed(time.Now().UTC().UnixNano())

	err := db.PostgresConnect()
	if err != nil{
		fmt.Println(err.Error())
	}
	defer db.DB.Close()

	httpd.StartServer()
}