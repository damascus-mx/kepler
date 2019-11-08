package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_app "github.com/damascus-mx/kepler/src/bin"
)

func main() {
	// Init root router
	/*
	*	TODO: Add PosgreSQL (pg) auth with driver
	 */

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	s3Bucket := os.Getenv("AWS_S3_ACCESS_KEY")
	fmt.Println(s3Bucket)

	http.ListenAndServe(": 3000", _app.InitApplication())
}
