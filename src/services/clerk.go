package services

import (
	"net/http"
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func CreateClerkClient() (clerk.Client, error) {
	clerkSecret := os.Getenv("CLERK_SECRET_KEY")

	clerkClient, err := clerk.NewClient(clerkSecret)

	if err != nil {
		return nil, err
	}

	return clerkClient, nil
}

func GetCurrentUser(r *http.Request) *clerk.User {
	ctx := r.Context()

	clerkClient, _ := CreateClerkClient()

	// ignore the ok and error because a user will be there as of require session middleware will return 403 else
	sessionClaims, _ := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
	user, _ := clerkClient.Users().Read(sessionClaims.Claims.Subject)

	return user
}

func ClerkActiveSession() (func(http.Handler) http.Handler, error) {
	clerkClient, err := CreateClerkClient()

	if err != nil {
		return nil, err
	}

	injectActiveSession := clerk.RequireSessionV2(clerkClient)

	return injectActiveSession, nil
}

func GetUser(userId string) (*clerk.User, error) {
	clerkClient, _ := CreateClerkClient()

	user, err := clerkClient.Users().Read(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsersByEmails(emails []string) ([]clerk.User, error) {
	clerkClient, _ := CreateClerkClient()

	users, err := clerkClient.Users().ListAll(clerk.ListAllUsersParams{EmailAddresses: emails})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUserPayment(userId string, stripeCustomerId string) {
	clerkClient, _ := CreateClerkClient()

	clerkClient.Users().Update(userId, &clerk.UpdateUser{PrivateMetadata: map[string]interface{}{"stripe_customer_id": stripeCustomerId}})

}
