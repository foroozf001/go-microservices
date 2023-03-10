package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const listenPort = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	connection, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer connection.Close()

	app := Config{
		Rabbit: connection,
	}

	log.Printf("broker service listening on port %s\n", listenPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", listenPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	log.Println(err)
}

func connect() (*amqp.Connection, error) {
	var retries int64
	var timeout = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("rabbitmq not ready..")
			retries++
		} else {
			log.Println("connected to rabbitmq")
			connection = c
			break
		}

		if retries > 5 {
			fmt.Println(err)
			return nil, err
		}

		timeout = time.Duration(math.Pow(float64(retries), 2)) * time.Second
		log.Println("timeout")
		time.Sleep(timeout)
		continue
	}

	return connection, nil
}
