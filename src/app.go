package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
	"log"
	"net/http"
	"os"
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
