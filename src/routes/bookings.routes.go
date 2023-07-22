package routes

import (
	"home-booking-api/src/controllers"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterBookingsRoutes(r *mux.Router, db *sqlx.DB) {
	r.Handle("/bookings", controllers.CreateBooking(db)).Methods("POST")
	r.Handle("/bookings/{bookingId}", controllers.ApproveBooking(db)).Methods("PUT")
}