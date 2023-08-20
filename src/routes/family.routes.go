package routes

import (
	"home-booking-api/src/controllers"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterFamilyRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	//r.Use(middleware.UserIsHouseAdmin(db, clerkClient))
	r.Handle("/family", controllers.CreateFamily(db)).Methods("POST")
	r.Handle("/family/{houseId}", controllers.FindFamilyForUser(db, clerkClient)).Methods("GET")
	r.Handle("/family/{familyId}", controllers.GetFamily(db)).Methods("GET")
	r.Handle("/family/{familyId}", controllers.RemoveFamily(db)).Methods("DELETE")
	r.Handle("/family/{familyId}", controllers.UpdateFamily(db)).Methods("PUT")
	r.Handle("/families/{houseId}", controllers.GetFamilies(db)).Methods("GET")
}
