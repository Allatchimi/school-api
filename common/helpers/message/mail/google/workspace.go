package googleMailHelper

import (
	"api/services/user/user/model"
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	googleAdmin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

// CreateGoogleWorkspaceUser creates a new user account in Google Workspace
func CreateGoogleWorkspaceUser(
	ctx *context.Context,
	credentials string,
	admin *model.User,
	user *model.User,
	password string,
) error {
	// Required OAuth scopes to manage users in Google Workspace
	scopes := []string{"https://www.googleapis.com/auth/admin.directory.user"}

	// Load credentials
	creds, err := google.JWTConfigFromJSON(
		[]byte(credentials),
		scopes...,
	)
	if err != nil {
		errMsg := "Failed to load credentials!"
		return fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
	}

	// Use domain-wide delegation to act on behalf of the admin
	creds.Subject = admin.Email

	newUser := &googleAdmin.User{
		PrimaryEmail: user.Email,
		Password:     password,
		Name: &googleAdmin.UserName{
			GivenName:  user.Info.FirstName,
			FamilyName: user.Info.LastName,
		},
		Gender: &googleAdmin.UserGender{
			Type: user.Info.Gender, // "male", "female"
		},
		Languages: []*googleAdmin.UserLanguage{
			{LanguageCode: "fr"},
			{LanguageCode: "en"},
		}, // Language codes must follow ISO 639-1 (e.g. "fr", "en")
	}

	// Initialize the Admin SDK Directory service
	srv, err := googleAdmin.NewService(*ctx, option.WithTokenSource(creds.TokenSource(*ctx)))
	if err != nil {
		errMsg := "Failed to create admin service!"
		return fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
	}

	// Create the user in Google Workspace
	_, err = srv.Users.Insert(newUser).Do()
	if err != nil {
		errMsg := "Failed to create Google account!"
		return fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
	}
	return nil
}
