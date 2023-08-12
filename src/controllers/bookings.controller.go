package controllers

import (
	"database/sql"
	"encoding/json"
	"home-booking-api/src/db/queries"
	"home-booking-api/src/models"
	"log"
	"net/http"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func CreateBooking(db *sqlx.DB, clerkClient clerk.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		sessionClaims, _ := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
		user, _ := clerkClient.Users().Read(sessionClaims.Claims.Subject)

		var createBookingPayload models.BookingModel

		createBookingPayload.UserBooking = user.ID

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

type Booking struct {
	User    clerk.User `json:"user"`
	Booking models.BookingModel `json:"booking"`
}

func GetHouseBookings(db *sqlx.DB, clerkClient clerk.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		houseId, ok := vars["houseId"]

		if !ok {
			http.Error(w, "no booking id provided", http.StatusBadRequest)
			return
		}

		var allBookings []Booking

		rows, err := db.Queryx(queries.GetBookingsForHouse, houseId)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "no bookings found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for rows.Next() {
			var bookingPayload Booking
			err := rows.StructScan(&bookingPayload.Booking)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			user, err := clerkClient.Users().Read(bookingPayload.Booking.UserBooking)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			bookingPayload.User = *user;
			log.Print(bookingPayload)

			allBookings = append(allBookings, bookingPayload)
		}

		defer rows.Close()

		json.NewEncoder(w).Encode(allBookings)

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
