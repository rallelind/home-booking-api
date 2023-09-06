package routes

import (
	"home-booking-api/src/controllers"
	"home-booking-api/src/middleware"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterHouseRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Handle("", controllers.GetUserHouses(db, clerkClient)).Methods("Get")
	r.Handle("", controllers.CreateHouse(db)).Methods("POST")
	r.Handle("/{houseId}", controllers.GetHouse(db)).Methods("GET")
	r.Handle("/{houseId}", controllers.RemoveHouse(db)).Methods("DELETE")
	r.Handle("/{houseId}/images", controllers.UploadHouseImages(db)).Methods("POST")
	r.Handle("/{houseId}/admin/approval", controllers.UpdateAdminNeedsToApprove(db)).Methods("PUT")
	r.Handle("/{houseId}/admin/users", controllers.UpdateAdminUsers(db)).Methods("PUT")

	r.Use(middleware.UserIsHouseAdmin(db, clerkClient))
	r.Handle("/families/{houseId}", controllers.CreateFamily(db)).Methods("POST")
	r.Handle("/families/{houseId}", controllers.GetFamilies(db, clerkClient)).Methods("GET")
}
