package controllers

import (
	"encoding/json"
	"home-booking-api/src/services"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func CreatePaymentCardSession(clerkClient clerk.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := services.GetCurrentUser(clerkClient, r)

		var payload struct {
			RedirectUrl string `json:"redirectUrl"`
		}

		err := json.NewDecoder(r.Body).Decode(&payload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		session, err := services.CreateCheckoutSession(user.EmailAddresses[0].EmailAddress, payload.RedirectUrl)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(session.URL)
	}
}
