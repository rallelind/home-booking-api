package routes

import (
	"home-booking-api/src/controllers"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterPaymentRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Handle("/webhook", controllers.WebhookHandler()).Methods("POST")
	r.Handle("/card/session", controllers.CreatePaymentCardSession(clerkClient)).Methods("POST")
	r.Handle("/methods", controllers.GetUserPaymentMethods(clerkClient)).Methods("GET")
	r.Handle("/methods/{paymentMethodId}", controllers.DeletePaymentMethod(clerkClient)).Methods("DELETE")
	r.Handle("/methods/{paymentMethodId}", controllers.SetPrimaryPaymentMethod(clerkClient)).Methods("PUT")
}
