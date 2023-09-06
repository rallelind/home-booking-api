package routes

import (
	"home-booking-api/src/controllers"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Handle("/family/{houseId}", controllers.FindFamilyForUser(db, clerkClient)).Methods("GET")
}
