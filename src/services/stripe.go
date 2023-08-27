package services

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"github.com/stripe/stripe-go/v75/customer"
	"github.com/stripe/stripe-go/v75/paymentmethod"
)

func CreateCustomer(email string) (*stripe.Customer, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	customer, err := customer.New(params)

	return customer, err
}

func CreateCheckoutSession(stripeCustomerId string, redirectUrl string) (*stripe.CheckoutSession, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSetup)),
		Customer:   stripe.String(stripeCustomerId),
		SuccessURL: stripe.String(fmt.Sprintf("%s?session_id={CHECKOUT_SESSION_ID}", redirectUrl)),
		CancelURL:  stripe.String(redirectUrl),
	}

	session, err := session.New(params)

	return session, err
}

func GetPaymentMethods(stripeCustomerId string) *[]stripe.PaymentMethod {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(stripeCustomerId),
		Type:     stripe.String("card"),
	}

	params.AddExpand("data.customer.invoice_settings")

	var paymentMethods []stripe.PaymentMethod

	i := paymentmethod.List(params)
	for i.Next() {
		pm := i.PaymentMethod()
		paymentMethods = append(paymentMethods, *pm)
	}

	return &paymentMethods
}

func SetPrimaryPaymentMethod(stripeCustomerId string, paymentMethodId string) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CustomerParams{}

	params.InvoiceSettings = &stripe.CustomerInvoiceSettingsParams{DefaultPaymentMethod: stripe.String(paymentMethodId)}

	customer.Update(stripeCustomerId, params)
}

func DeletePaymentMethod(paymentMethodId string) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	paymentmethod.Detach(paymentMethodId, nil)
}
