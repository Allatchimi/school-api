package googleMailHelper

import (
	"api/common/utils"
	"api/services/user/user/model"
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	googleAdmin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

// CreateGoogleWorkspaceUser creates a new user account in Google Workspace
func CreateGoogleWorkspaceUser(
	ctx context.Context,
	credentialsJSON string,
	adminEmail string,
	user *model.User,
	password string,
) (*googleAdmin.User, error) {
	// Validate inputs early
	if user == nil || user.Info == nil || !utils.IsEmailValid(user.Email) ||
		len(password) == 0 || len(credentialsJSON) == 0 {
		return nil, fmt.Errorf("invalid inputs")
	}

	// Required scopes to manage users; add more if needed (e.g., user.security)
	scopes := []string{
		"https://www.googleapis.com/auth/admin.directory.user",
	}

	// Load service account credentials
	creds, err := google.JWTConfigFromJSON([]byte(credentialsJSON), scopes...)
	if err != nil {
		return nil, fmt.Errorf("failed to load credentials: %w", err)
	}

	// Domain-Wide Delegation - act on behalf of a super admin in the Workspace domain
	// IMPORTANT: Replace with a real super admin email
	if !utils.IsEmailValid(adminEmail) {
		return nil, fmt.Errorf("missing GOOGLE_WORKSPACE_ADMIN_EMAIL (super admin email for domain-wide delegation)")
	}
	creds.Subject = adminEmail

	// Build the user payload; keep optional fields only if valid
	var gender *googleAdmin.UserGender
	if user.Info.Gender == "male" || user.Info.Gender == "female" {
		gender = &googleAdmin.UserGender{
			Type: user.Info.Gender,
		}
	}
	newUser := &googleAdmin.User{
		PrimaryEmail: user.Email,
		Password:     password, // Plaintext allowed; Google stores it securely. Ensure it matches domain password policy.
		Name: &googleAdmin.UserName{
			GivenName:  user.Info.FirstName,
			FamilyName: user.Info.LastName,
			// FullName is optional; Admin API can compute display names; keep consistent with your needs
			FullName: fmt.Sprintf("%s %s", user.Info.FirstName, user.Info.LastName),
		},
		// Gender is optional; ensure the value matches allowed enum; if unsure, omit
		Gender: gender,
		// Languages must be ISO 639-1 codes; optional
		Languages: []*googleAdmin.UserLanguage{
			{LanguageCode: "fr"},
			{LanguageCode: "en"},
		},
		// Optional orgUnitPath, change password at next login, etc.
		// OrgUnitPath: "/",
		// ChangePasswordAtNextLogin: true,
	}

	// Initialize Admin SDK Directory service using the token source from service account
	srv, err := googleAdmin.NewService(ctx, option.WithTokenSource(creds.TokenSource(ctx)))
	if err != nil {
		return nil, fmt.Errorf("failed to create admin service: %w", err)
	}

	// Create the user
	result, err := srv.Users.Insert(newUser).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to create Google account: %w", err)
	}
	return result, nil
}
