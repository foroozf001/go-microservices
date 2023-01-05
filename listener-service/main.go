package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer connection.Close()

	log.Println("listening for rabbitmq messages..")

	consumer, err := event.NewConsumer(connection)
	if err != nil {
		panic(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
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
