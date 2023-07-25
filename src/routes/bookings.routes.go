package routes

import (
	"home-booking-api/src/controllers"
	"home-booking-api/src/middleware"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterBookingsRoutes(r *mux.Router, db *sqlx.DB) {
	r.Use(middleware.UserPartOfHouse(db))
	r.Handle("/bookings", controllers.CreateBooking(db)).Methods("POST")
	r.Handle("/bookings/{bookingId}", controllers.ApproveBooking(db)).Methods("PUT")
	r.Handle("/bookings/{bookingId}", controllers.RemoveBooking(db)).Methods("DELETE")
}