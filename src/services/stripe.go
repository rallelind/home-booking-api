package services

import (
	"os"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/customer"
)

func CreateCustomer(email string) (*stripe.Customer, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	customer, err := customer.New(params)

	return customer, err
}
