package services

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
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

func CreateCheckoutSession(email string, redirectUrl string) (*stripe.CheckoutSession, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	customer, err := CreateCustomer(email)

	if err != nil {
		return nil, err
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSetup)),
		Customer:   stripe.String(customer.ID),
		SuccessURL: stripe.String(fmt.Sprintf("%s?session_id={CHECKOUT_SESSION_ID}", redirectUrl)),
		CancelURL:  stripe.String(redirectUrl),
	}

	session, err := session.New(params)

	return session, err
}
