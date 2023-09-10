package routes

import (
	"home-booking-api/src/controllers"
	"home-booking-api/src/middleware"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterBookingsRoutes(r *mux.Router, db *sqlx.DB) {
	r.Use(middleware.UserPartOfHouse(db))
	r.Handle("/{houseId}", controllers.CreateBooking(db)).Methods("POST")
	r.Handle("/{houseId}/today", controllers.GetTodayBooking(db)).Methods("GET")
	r.Handle("/{houseId}/past", controllers.GetPastBookings(db)).Methods("GET")
	r.Handle("/{houseId}/bookings", controllers.GetHouseBookings(db)).Methods("GET")
	r.Handle("/{houseId}/bookings/{bookingId}", controllers.ApproveBooking(db)).Methods("PUT")
	r.Handle("/{houseId}/bookings/{bookingId}", controllers.RemoveBooking(db)).Methods("DELETE")
}
