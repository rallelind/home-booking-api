package routes

import (
	"home-booking-api/src/controllers"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterPaymentRoutes(r *mux.Router, db *sqlx.DB, clerkClient clerk.Client) {
	r.Handle("/payment/webhook", controllers.WebhookHandler()).Methods("POST")
	r.Handle("/payment/card/session", controllers.CreatePaymentCardSession(clerkClient)).Methods("POST")
	r.Handle("/payment/methods", controllers.GetUserPaymentMethods(clerkClient)).Methods("GET")
	r.Handle("/payment/methods/{paymentMethodId}", controllers.DeletePaymentMethod(clerkClient)).Methods("DELETE")
	r.Handle("/payment/methods/{paymentMethodId}", controllers.SetPrimaryPaymentMethod(clerkClient)).Methods("PUT")
}
