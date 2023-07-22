package controllers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func CreateBooking(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}