package routes

import (
	"home-booking-api/src/controllers"
	"home-booking-api/src/middleware"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterBookingsRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Use(middleware.UserPartOfHouse(db, clerkClient))
	r.Handle("", controllers.CreateBooking(db, clerkClient)).Methods("POST")
	r.Handle("/{houseId}/today", controllers.GetTodayBooking(db, clerkClient)).Methods("GET")
	r.Handle("/{houseId}/past", controllers.GetPastBookings(db, clerkClient)).Methods("GET")
	r.Handle("/bookings/{houseId}", controllers.GetHouseBookings(db, clerkClient)).Methods("GET")
	r.Handle("/bookings/{bookingId}", controllers.ApproveBooking(db)).Methods("PUT")
	r.Handle("/bookings/{bookingId}", controllers.RemoveBooking(db)).Methods("DELETE")
}
