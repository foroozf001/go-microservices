package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const listenPort = "80"

var retries int64 = 0

const maxRetries = 10

const sleep = 5 * time.Second

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Printf("authentication service listening on port %s\n", listenPort)

	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to postgres")
	}

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", listenPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	log.Panic(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet available..")
			retries++
		} else {
			log.Println("connected to postgres")
			return conn
		}

		if retries > maxRetries {
			log.Println(err)
			return nil
		}

		log.Println("retrying postgres connection..")
		time.Sleep(sleep)
		continue
	}
}
