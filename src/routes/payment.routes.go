package routes

import (
	"home-booking-api/src/controllers"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterPaymentRoutes(r *mux.Router, db *sqlx.DB) {
	r.Handle("/webhook", controllers.WebhookHandler()).Methods("POST")
	r.Handle("/card/session", controllers.CreatePaymentCardSession()).Methods("POST")
	r.Handle("/methods", controllers.GetUserPaymentMethods()).Methods("GET")
	r.Handle("/methods/{paymentMethodId}", controllers.DeletePaymentMethod()).Methods("DELETE")
	r.Handle("/methods/{paymentMethodId}", controllers.SetPrimaryPaymentMethod()).Methods("PUT")
}
