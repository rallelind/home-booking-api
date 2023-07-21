package routes

import (
	"home-booking-api/src/controllers"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterHouseRoutes(r *mux.Router, db *sqlx.DB) {
	r.Handle("/house", controllers.CreateHouse(db)).Methods("POST")
	r.Handle("/house/{houseId}", controllers.UpdateHouse(db)).Methods("PUT")
}
