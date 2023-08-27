package controllers

import (
	"encoding/json"
	"fmt"
	"home-booking-api/src/services"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/customer"
	"github.com/stripe/stripe-go/v75/setupintent"
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

func WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const MaxBodyBytes = int64(65536)
		r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
		payload, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

		event := stripe.Event{}

		if err := json.Unmarshal(payload, &event); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch event.Type {
		case "checkout.session.completed":
			var checkoutSession stripe.CheckoutSession
			err := json.Unmarshal(event.Data.Raw, &checkoutSession)

			if err != nil {
				log.Printf("Error parsing webhook JSON: %v\n", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Print(checkoutSession)

			result, err := setupintent.Get(checkoutSession.SetupIntent.ID, nil)

			if err != nil {
				log.Printf("Error parsing webhook JSON: %v\n", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			params := &stripe.CustomerParams{InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{DefaultPaymentMethod: stripe.String(result.PaymentMethod.ID)}}

			customer.Update(result.Customer.ID, params)
		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		}
		w.WriteHeader(http.StatusOK)

	}

}
