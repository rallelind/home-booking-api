package middleware

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func UserIsHouseAdmin(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
