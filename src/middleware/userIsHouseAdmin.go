package middleware

import (
	"database/sql"
	"home-booking-api/src/db/queries"
	"home-booking-api/src/models"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/jmoiron/sqlx"
)

func UserIsHouseAdmin(db *sqlx.DB, clerkClient clerk.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			sessionClaims, _ := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
			user, _ := clerkClient.Users().Read(sessionClaims.Claims.Subject)

			var house models.HouseModel

			err := db.QueryRowx(queries.UserHouseAdminQuery, user.EmailAddresses[0].EmailAddress).StructScan(&house)

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
