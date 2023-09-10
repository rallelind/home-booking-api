package middleware

import (
	"database/sql"
	"home-booking-api/src/db/queries"
	"home-booking-api/src/models"
	"home-booking-api/src/services"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func UserPartOfHouse(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r)
			id, ok := vars["houseId"]

			if !ok {
				http.Error(w, "please provide the house id", http.StatusBadRequest)
				return
			}

			user := services.GetCurrentUser(r)

			var family models.FamilyModel

			err := db.QueryRowx(queries.UserPartOfHouseQuery, user.EmailAddresses[0].EmailAddress, id).StructScan(&family)

			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, "You are not allowed as you are not admin", http.StatusForbidden)
					return
				}

				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
