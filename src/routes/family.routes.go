package routes

import (
	"home-booking-api/src/controllers"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterFamilyRoutes(r *mux.Router, db *sqlx.DB) {
	r.Handle("/family", controllers.CreateFamily(db)).Methods("POST")
	r.Handle("/family/{familyId}", controllers.GetFamily(db)).Methods("GET")
	r.Handle("/family/{familyId}", controllers.RemoveFamily(db)).Methods("DELETE")
	r.Handle("/family/{familyId}", controllers.UpdateFamily(db)).Methods("PUT")
}
