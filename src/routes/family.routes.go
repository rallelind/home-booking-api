package routes

import (
	"home-booking-api/src/controllers"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterFamilyRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Handle("/{familyId}", controllers.GetFamily(db)).Methods("GET")
	r.Handle("/{familyId}", controllers.RemoveFamily(db)).Methods("DELETE")
	r.Handle("/{familyId}", controllers.UpdateFamily(db)).Methods("PUT")
	r.Handle("/{familyId}/cover-image", controllers.UpdateFamilyCoverImage(db)).Methods("PUT")
}
