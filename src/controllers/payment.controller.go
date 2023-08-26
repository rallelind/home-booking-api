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

		customerId, ok := user.PrivateMetadata.(map[string]interface{})["stripe_customer_id"].(string)

		if !ok {
			customer, err := services.CreateCustomer(user.EmailAddresses[0].EmailAddress)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			clerkClient.Users().Update(user.ID, &clerk.UpdateUser{PrivateMetadata: map[string]interface{}{"stripe_customer_id": customer.ID}})
			customerId = customer.ID
		}

		session, err := services.CreateCheckoutSession(customerId, payload.RedirectUrl)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(session.URL)
	}
}
