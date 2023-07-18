package app

import (
	"log"
	"net/http"
	"database/sql"
	"os"
	"github.com/lpernett/godotenv"
)

func App() {

	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	defer db.Close()

	mux := http.NewServeMux()

	http.ListenAndServe(":8080", mux)

}
