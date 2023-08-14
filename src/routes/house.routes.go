package routes

import (
	"home-booking-api/src/controllers"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterHouseRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Handle("/house", controllers.CreateHouse(db)).Methods("POST")
	r.Handle("/house", controllers.GetUserHouses(db, clerkClient)).Methods("Get")
	r.Handle("/house/{houseId}", controllers.UpdateHouse(db)).Methods("PUT")
	r.Handle("/house/{houseId}", controllers.GetHouse(db)).Methods("GET")
	r.Handle("/house/{houseId}", controllers.RemoveHouse(db)).Methods("DELETE")
	r.Handle("/house/{houseId}/images", controllers.UploadHouseImages(db)).Methods("POST")
}
