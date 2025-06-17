package mail

import (
	"fmt"
	"strings"
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
