package helpers

import (
	"fmt"
	"strings"
	"time"
)

// GenerateEmailFromFullName takes the first name and last name,
// extracts the first word from each (in case there are multiple),
// converts them to lowercase, and formats an email address.
func GenerateEmailFromFullName(firstName string, lastName string, emailDomain string) string {
	// Split the names into words by spaces
	firstNameParts := strings.Fields(firstName)
	lastNameParts := strings.Fields(lastName)

	// Take the first word from each
	first := ""
	last := ""
	if len(firstNameParts) > 0 {
		first = firstNameParts[0]
	}
	if len(lastNameParts) > 0 {
		last = lastNameParts[0]
	}

	// Construct the email address in the format: firstname.lastname@domain
	email := fmt.Sprintf("%s.%s@%s", strings.ToLower(first), strings.ToLower(last), strings.ToLower(emailDomain))
	return email
}

// GeneratePasswordFromUserInfo generates a password from user information.
func GeneratePasswordFromUser(firstName string, lastName string, birthday *time.Time) string {
	var formattedBirthday = time.Now().Year()
	firstNameParts := strings.Split(firstName, " ")
	if len(firstNameParts) > 0 {
		firstName = firstNameParts[0]
	}
	lastNameParts := strings.Split(lastName, " ")
	if len(lastNameParts) > 0 {
		lastName = lastNameParts[0]
	}
	if birthday != nil {
		formattedBirthday = birthday.Year()
	}
	var password string = ""
	if len(firstName) > 0 && len(lastName) > 0 && formattedBirthday > 0 {
		// Lowercase
		firstName = strings.ToLower(firstName)
		lastName = strings.ToLower(lastName)
		// Format
		password = fmt.Sprintf("%s%s%d", firstName, lastName, formattedBirthday)
	}
	return password
}
