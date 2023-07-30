package controllers

import (
	"encoding/json"
	"home-booking-api/src/db/queries"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type BookingPayload struct {
	Id        int       `db:"id" json:"id"`
	StartDate time.Time `db:"start_date" json:"start_date"`
	EndDate   time.Time `db:"end_date" json:"end_date"`
	Approved  bool      `db:"approved" json:"approved"`
	HouseId   int       `db:"house_id" json:"house_id"`
}

func CreateBooking(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// implement middleware to make sure the user is part of house

		w.Header().Set("Content-Type", "application/json")

		var createBookingPayload BookingPayload

		err := json.NewDecoder(r.Body).Decode(&createBookingPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if createBookingPayload.EndDate == createBookingPayload.StartDate {
			http.Error(w, "dates can't be the same", http.StatusBadRequest)
			return
		}

		var count int

		err = db.QueryRow(queries.FindBookingForSpecificDateQuery, &createBookingPayload.StartDate, &createBookingPayload.EndDate).Scan(&count)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if count > 0 {
			http.Error(w, "The chosen dates is overlapping with other dates", http.StatusBadRequest)
			return
		}

		_, err = db.NamedExec(queries.CreateBookingQuery, createBookingPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("successfully created")
	}
}

func RemoveBooking(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		bookingId, ok := vars["bookingId"]

		if !ok {
			http.Error(w, "no booking id provided", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(queries.DeleteBookingQuery, bookingId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("successfully deleted")
	}
}

type BookingApprovalPayload struct {
	Approved bool `db:"approved" json:"approved"`
}

func ApproveBooking(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		bookingId, ok := vars["bookingId"]

		if !ok {
			http.Error(w, "please provide booking id", http.StatusBadRequest)
			return
		}

		var bookingApprovalPayload BookingApprovalPayload

		err := json.NewDecoder(r.Body).Decode(&bookingApprovalPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var updatedStartDate time.Time
		var updatedEndDate time.Time
		var updatedApproved bool
		var updatedId int

		err = db.QueryRow(queries.UpdateBookingQuery, bookingApprovalPayload.Approved, bookingId).Scan(&updatedStartDate, &updatedEndDate, &updatedApproved, &updatedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !updatedApproved {
			json.NewEncoder(w).Encode("updated")
			return
		}

		_, err = db.Exec(queries.DeleteOverlappingBookingsQuery, updatedStartDate, updatedEndDate, updatedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("updated")

	}
}
