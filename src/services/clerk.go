package services

import (
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func GetCurrentUser(clerkClient clerk.Client, r *http.Request) *clerk.User {
	ctx := r.Context()

	// ignore the ok and error because a user will be there as of require session middleware will return 403 else
	sessionClaims, _ := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
	user, _ := clerkClient.Users().Read(sessionClaims.Claims.Subject)

	return user
}
