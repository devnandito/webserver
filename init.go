package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"web_server/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initDB() {
	load := godotenv.Load("envs/.env")

	if load != nil {
		panic(load)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PWD := os.Getenv("DB_PWD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", DB_PORT, DB_HOST, DB_USER, DB_PWD, DB_NAME)

	var err error

	models.DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
}