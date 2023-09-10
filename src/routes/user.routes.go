package routes

import (
	"home-booking-api/src/controllers"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(r *mux.Router, db *sqlx.DB) {
	r.Handle("/family/{houseId}", controllers.FindFamilyForUser(db)).Methods("GET")
}
