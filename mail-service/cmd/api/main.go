package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Mailer Mail
}

const listenPort = "80"

func main() {
	app := Config{
		Mailer: createMail(),
	}

	log.Printf("mail service listening on port %s\n", listenPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", listenPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	log.Println(err)
}

func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Ecryption:   os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
	}

	return m
}
