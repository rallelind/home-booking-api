package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

		var count int

		sqlQueryExistingBookings := `
			SELECT COUNT(*) FROM bookings
			WHERE (start_date, end_date) OVERLAPS ($1, $2)
		`

		err = db.QueryRow(sqlQueryExistingBookings, &createBookingPayload.StartDate, &createBookingPayload.EndDate).Scan(&count)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if count > 0 {
			http.Error(w, "The chosen dates is overlapping with other dates", http.StatusBadRequest)
			return
		}

		sqlInsertBooking := `
			INSERT INTO bookings (start_date, end_date, approved, house_id)
			VALUES (:start_date, :end_date,
				CASE
					WHEN (SELECT admin_needs_to_approve FROM houses WHERE id = :house_id) = TRUE THEN FALSE
					ELSE TRUE
				END,
			:house_id
			)
		`

		_, err = db.NamedExec(sqlInsertBooking, createBookingPayload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("successfully created")
	}
}
