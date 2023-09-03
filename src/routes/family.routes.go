package routes

import (
	"home-booking-api/src/controllers"
	"home-booking-api/src/middleware"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterFamilyRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Use(middleware.UserIsHouseAdmin(db, clerkClient))
	r.Handle("/", controllers.CreateFamily(db)).Methods("POST")
	r.Handle("/me/{houseId}", controllers.FindFamilyForUser(db, clerkClient)).Methods("GET")
	r.Handle("/{familyId}", controllers.GetFamily(db)).Methods("GET")
	r.Handle("/{familyId}", controllers.RemoveFamily(db)).Methods("DELETE")
	r.Handle("/{familyId}", controllers.UpdateFamily(db)).Methods("PUT")
	r.Handle("/families/{houseId}", controllers.GetFamilies(db, clerkClient)).Methods("GET")
	r.Handle("/{familyId}/cover-image", controllers.UpdateFamilyCoverImage(db)).Methods("PUT")
}
