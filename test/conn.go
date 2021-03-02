package main

import (
	"context"
	"fmt"
	"flag"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var pool *sql.DB

func IntiDB() {
	load := godotenv.Load(".env")
	
	if load != nil {
		panic(load)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PWD := os.Getenv("DB_PWD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	pg_conn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", DB_PORT, DB_HOST, DB_USER, DB_PWD, DB_NAME)

	id := flag.Int64("id", 0, "person ID to find")
	dsn := flag.String("dsn", pg_conn, "connection data source name")

	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	if *id == 0 {
		log.Fatal("missing person ID")
	}
	var err error

	pool, err := sql.Open("postgres", *dsn)
	
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()
	Ping(ctx)
	Query(ctx, *id)
}

func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func Query(ctx context.Context, id int64) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var table4 Table4
	query := "SELECT email, fullname FROM table4 WHERE pk = $1"
	err := pool.QueryRow(query, 1).Scan(&table4.Email, &table4.Fullname)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("Name: ", &table4.Fullname)
}
