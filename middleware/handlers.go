package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	// check error load environment
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Open connection to database
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// Check connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database successfully")

	//return the connection
	return db
}
