package app

import (
	"home-booking-api/src/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

func App() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	defer db.Close()

	mux := mux.NewRouter()
	mux.Handle("/house", controllers.CreateHouse(db)).Methods("POST")

	http.ListenAndServe(":8080", mux)

}
