package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Uchencho/serverlessTest/internal"
	"github.com/joho/godotenv"
)

const defaultServerAddress = "127.0.0.1:4000"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found, with error: %s", err)
	}
}

func main() {

	a := internal.New()

	log.Println(fmt.Sprintf("Starting server on address:%s", defaultServerAddress))
	log.Fatal(http.ListenAndServe(defaultServerAddress, a.Handler()))
}
